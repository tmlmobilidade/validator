package tests

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/vehicles/validations"
	"testing"
)

func TestParseVehicles_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	raw := types.VehicleRaw{
		VehicleId:         "V1",
		AgencyId:          "A1",
		LicensePlate:      "AB-12-CD",
		Make:              "Make",
		Model:             "Model",
		Owner:             "Owner",
		RegistrationDate:  "2021-01-01",
		AvailableSeats:    "10",
		AvailableStanding: "10",
		Typology:          "1.0",
		Propulsion:        "1",
		Emission:          "1",
		Climatization:     "1",
		Wheelchair:        "1",
		LoweredFloor:      "1",
		Ramp:              "1",
		Kneeling:          "1",
		StaticInformation: "1",
		OnboardMonitor:    "1",
		FrontDisplay:      "1",
		RearDisplay:       "1",
		SideDisplay:       "1",
		InternalSound:     "1",
		ExternalSound:     "1",
		ConsumptionMeter:  "1",
		Bicycles:          "1",
		PassengerCounting: "1",
		VideoSurveillance: "1",
	}

	validations.ParseVehicles(raw, row)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "ParseVehicles_ValidInput", types.SEVERITY_ERROR)
}

func TestParseVehicles_InvalidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	raw := types.VehicleRaw{
		VehicleId:         "V1",
		AgencyId:          "A1",
		LicensePlate:      "not_a_license_plate",
		Make:              "not_a_make",
		Model:             "not_a_model",
		Owner:             "not_a_owner",
		RegistrationDate:  "not_a_registration_date",
		AvailableSeats:    "not_a_number",
		AvailableStanding: "not_a_number",
		Typology:          "not_a_typology",
		Propulsion:        "not_a_propulsion",
		Emission:          "not_a_emission",
		Climatization:     "not_a_climatization",
		Wheelchair:        "not_a_wheelchair",
		LoweredFloor:      "not_a_lowered_floor",
		Ramp:              "not_a_ramp",
		Kneeling:          "not_a_kneeling",
		StaticInformation: "not_a_static_information",
		OnboardMonitor:    "not_a_onboard_monitor",
		FrontDisplay:      "not_a_front_display",
		RearDisplay:       "not_a_rear_display",
		SideDisplay:       "not_a_side_display",
		InternalSound:     "not_a_internal_sound",
		ExternalSound:     "not_a_external_sound",
		ConsumptionMeter:  "not_a_consumption_meter",
		Bicycles:          "not_a_bicycles",
		PassengerCounting: "not_a_passenger_counting",
		VideoSurveillance: "not_a_video_surveillance",
	}
	validations.ParseVehicles(raw, row)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "ParseVehicles_InvalidInput", types.SEVERITY_ERROR)
}
