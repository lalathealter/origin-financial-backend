package controllers

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestCalculateRiskProfile(t *testing.T) {

	for cinfo, rpfile := range testConvMap {
		got := calculateRiskProfile(&cinfo)
		assert.Equal(t, got, rpfile, fmt.Sprintf("%v", cinfo))
	}

}

var testConvMap = map[ClientInformationRisks]RiskProfile{
	mockCinfoInit: {
		Auto:       Regular,
		Disability: Ineligible,
		Home:       Economic,
		Life:       Regular,
	},
	{
		Age:           61,
		Dependents:    0,
		Income:        250000,
		Marital:       Married,
		RiskQuestions: [...]uint{0, 0, 0},
		House:         HouseData{NoHouse},
		Vehicle:       VehicleData{0},
	}: {
		Home:       Ineligible,
		Auto:       Ineligible,
		Life:       Ineligible,
		Disability: Ineligible,
	},
	{
		Age:           29,
		Dependents:    1,
		Income:        0,
		Marital:       Married,
		RiskQuestions: [...]uint{0, 4, 4},
		House:         HouseData{Owned},
		Vehicle:       VehicleData{uint(time.Now().Year() - 5)},
	}: {
		Home:       Economic,
		Auto:       Regular,
		Life:       Regular,
		Disability: Ineligible,
	},
	{
		Age:           35,
		Dependents:    0,
		Marital:       Married,
		Income:        250000,
		RiskQuestions: [...]uint{1, 0, 1},
		House:         HouseData{NoHouse},
		Vehicle:       VehicleData{uint(time.Now().Year()) - 10},
	}: {
		Home:       Ineligible,
		Auto:       Economic,
		Life:       Regular,
		Disability: Economic,
	},
	{
		Age:           52,
		Dependents:    4,
		Income:        12500,
		Marital:       Married,
		RiskQuestions: [...]uint{1, 4, 0},
		House:         HouseData{Mortgaged},
		Vehicle:       VehicleData{0},
	}: {
		Auto:       Ineligible,
		Disability: Responsible,
		Home:       Responsible,
		Life:       Responsible,
	},
	{
		Age:           33,
		Dependents:    0,
		Income:        2500,
		Marital:       Single,
		RiskQuestions: [...]uint{1, 1, 1},
		House:         HouseData{Mortgaged},
		Vehicle:       VehicleData{uint(time.Now().Year()) - 2},
	}: {
		Home:       Responsible,
		Auto:       Responsible,
		Life:       Regular,
		Disability: Responsible,
	},
}
