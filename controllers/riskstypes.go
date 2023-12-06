package controllers

import "time"

type RiskScoreHolder map[RiskFactor]int

func (rsh *RiskScoreHolder) AddToEveryField(val int) {
	for key := range *rsh {
		(*rsh)[key] += val
	}
}

func (rsh RiskScoreHolder) SetIneligible(key RiskFactor) {
	delete(rsh, key)
}

func (rsh *RiskScoreHolder) AddScoreTo(key RiskFactor, val int) {
	_, ok := (*rsh)[key]
	if !ok {
		return
	}
	(*rsh)[key] += val
}

func (rsh *RiskScoreHolder) ConcludeFactorScore(key RiskFactor) RiskScoreText {
	score, ok := (*rsh)[key]
	if !ok {
		return Ineligible
	}

	switch {
	case score <= 0:
		return Economic
	case score >= 1 && score <= 2:
		return Regular
	}
	return Responsible
}

type RiskFactor string

const (
	Life       RiskFactor = "Life"
	Disability RiskFactor = "Disability"
	Home       RiskFactor = "Home"
	Auto       RiskFactor = "Auto"
)

type RiskProfile struct {
	Auto       RiskScoreText `json:"auto"`
	Disability RiskScoreText `json:"disability"`
	Home       RiskScoreText `json:"home"`
	Life       RiskScoreText `json:"life"`
}

type RiskScoreText string

const (
	Ineligible  RiskScoreText = "ineligible"
	Economic    RiskScoreText = "economic"
	Regular     RiskScoreText = "regular"
	Responsible RiskScoreText = "responsible"
)

type ClientInformationRisks struct {
	Age           uint
	Dependents    uint
	Income        uint
	Marital       MaritalStatus      `json:"marital_status" binding:"oneof='married' 'single',required"`
	RiskQuestions RiskQuestionsSlice `json:"risk_questions" binding:"required"`
	House         HouseData          `binding:"omitempty"`
	Vehicle       VehicleData        `binding:"omitempty"`
}

type RiskQuestionsSlice [3]uint

func (rqs *RiskQuestionsSlice) GetBaseRiskScore() int {
	score := 0
	for _, v := range rqs {
		if v > 0 {
			score++
		}
	}
	return score
}

type MaritalStatus string

const (
	Married MaritalStatus = "married"
	Single  MaritalStatus = "single"
)

type OwnershipStatus string

const (
	NoHouse   OwnershipStatus = ""
	Owned     OwnershipStatus = "owned"
	Mortgaged OwnershipStatus = "mortgaged"
)

type HouseData struct {
	Ownership OwnershipStatus `json:"ownership_status" binding:"oneof=owned mortgaged,required"`
}

type VehicleData struct {
	Year uint `binding:"required"`
}

func (vhd VehicleData) WasProducedLessThanYearsAgo(yearsAgo int) bool {
	return time.Now().Year()-int(vhd.Year) <= yearsAgo
}

func (cinfo ClientInformationRisks) HasNoVehicle() bool {
	return cinfo.Vehicle.Year == 0
}
