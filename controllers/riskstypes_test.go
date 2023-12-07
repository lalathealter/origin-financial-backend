package controllers

import (
	"reflect"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestVehicleData(t *testing.T) {

	veh := VehicleData{Year: uint(time.Now().Year())}

	assert.Equal(t, veh.WasProducedLessThanYearsAgo(5), true)
	veh.Year -= 5
	assert.Equal(t, veh.WasProducedLessThanYearsAgo(5), true)
	veh.Year -= 10
	assert.Equal(t, veh.WasProducedLessThanYearsAgo(2), false)
	veh.Year += 22
	assert.Equal(t, veh.WasProducedLessThanYearsAgo(4), true)
}

var mockCinfoInit = ClientInformationRisks{
	35, 2, 0, "married", [...]uint{0, 1, 0},
	HouseData{Ownership: Owned},
	VehicleData{Year: uint(time.Now().Year() - 5)},
}

func TestRiskProfileDefinition(t *testing.T) {
	rsh := MakeRiskScoreHolder(&mockCinfoInit)
	rpfile := concludeScoresIntoRiskProfile(rsh)

	fields := reflect.ValueOf(rpfile).NumField()
	assert.Equal(t, fields, len(RiskFactorsColl))
}

