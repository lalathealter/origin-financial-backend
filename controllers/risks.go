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

	c.JSON(http.StatusCreated, clientData)
}

type ClientInformationRisks struct {
	Age           uint               `binding:"gt=0,required"`
	Dependents    uint               `binding:"gte=0,required"`
	House         HouseData          `binding:"omitempty"`
	Income        uint               `binding:"gt=0,required"`
	Marital       MaritalStatus      `json:"marital_status" binding:"oneof='married' 'single',required"`
	RiskQuestions RiskQuestionsSlice `json:"risk_questions" binding:"required"`
	Vehicle       VehicleData        `binding:"omitempty"`
}

type RiskQuestionsSlice [3]uint

type MaritalStatus string

const (
	Married MaritalStatus = "married"
	Single  MaritalStatus = "single"
)

type OwnershipStatus string

const (
	Owned     OwnershipStatus = "owned"
	Mortgaged OwnershipStatus = "mortgaged"
)

type HouseData struct {
	Ownership OwnershipStatus `json:"ownership_status" binding:"oneof=owned mortgaged,required"`
}

type VehicleData struct {
	Year uint `binding:"gt=0,required"`
}
