package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRisksCalculation(c *gin.Context) {

	var clientData ClientInformationRisks
	if err := c.ShouldBindJSON(&clientData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	riskProf := calculateRiskProfile(&clientData)
	c.JSON(http.StatusCreated, riskProf)
}

func calculateRiskProfile(cinfo *ClientInformationRisks) RiskProfile {
	rsholder := MakeRiskScoreHolder(cinfo)

	if cinfo.Income == 0 {
		rsholder.SetIneligible(Disability)
	} else if cinfo.Income > 200_000 {
		rsholder.AddToEveryField(-1)
	}

	if cinfo.Age > 60 {
		rsholder.SetIneligible(Disability)
		rsholder.SetIneligible(Life)
	} else if cinfo.Age < 30 {
		rsholder.AddToEveryField(-2)
	} else if cinfo.Age >= 30 && cinfo.Age <= 40 {
		rsholder.AddToEveryField(-1)
	}

	if cinfo.House.Ownership == NoHouse {
		rsholder.SetIneligible(Home)
	} else if cinfo.House.Ownership == Mortgaged {
		rsholder.AddScoreTo(Disability, 1)
		rsholder.AddScoreTo(Home, 1)
	}

	if cinfo.Dependents > 0 {
		rsholder.AddScoreTo(Disability, 1)
		rsholder.AddScoreTo(Life, 1)
	}

	if cinfo.Marital == Married {
		rsholder.AddScoreTo(Life, 1)
		rsholder.AddScoreTo(Disability, -1)
	}

	if cinfo.HasNoVehicle() {
		rsholder.SetIneligible(Auto)
	} else if cinfo.Vehicle.WasProducedLessThanYearsAgo(5) {
		rsholder.AddScoreTo(Auto, 1)
	}

	return concludeScoresIntoRiskProfile(rsholder)
}

func concludeScoresIntoRiskProfile(rsh RiskScoreHolder) RiskProfile {
	rpfile := RiskProfile{}
	rpfile.Auto = rsh.ConcludeFactorScore(Auto)
	rpfile.Home = rsh.ConcludeFactorScore(Home)
	rpfile.Disability = rsh.ConcludeFactorScore(Disability)
	rpfile.Life = rsh.ConcludeFactorScore(Life)
	return rpfile
}

func MakeRiskScoreHolder(cinfo *ClientInformationRisks) RiskScoreHolder {
	rsholder := RiskScoreHolder{}
	baseScore := cinfo.RiskQuestions.GetBaseRiskScore()
	for _, factor := range RiskFactorsColl {
		rsholder[factor] = baseScore
	}
	return rsholder
}
