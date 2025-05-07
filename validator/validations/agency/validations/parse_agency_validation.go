package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseAgencyValidation(rawAgency map[string]string, row int, gtfs *types.Gtfs) (agency types.Agency) {
	message := types.Message{
		Field:   "",
		FileName: "agency.txt",
		Rows: []int{row},
		Message: "",
		Severity: types.SEVERITY_ERROR,
		ValidationID: "agency_parse",
	}

	var agencyName, agencyTimezone, agencyUrl string
	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone, agencyId string

	// Define field mappings
	fieldMappings := map[string]*string{
		"agency_name":     &agencyName,
		"agency_url":      &agencyUrl,
		"agency_timezone": &agencyTimezone,
		"agency_id":       &agencyId,
		"agency_lang":     &agencyLang,
		"agency_phone":    &agencyPhone,
		"agency_fare_url": &agencyFareUrl,
		"agency_email":    &agencyEmail,
	}

	// Loop through fields and parse each one
	for field, target := range fieldMappings {
		msg := lib.ParseStringToPrimitive(rawAgency[field], target)
		if msg != "" {
			message.Message = msg
			message.Field = field
			services.AppMessageService.AddMessage(message)
		}
	}

	agency.AgencyEmail = lib.IfThenElse(agencyEmail != "", &agencyEmail, nil)
	agency.AgencyFareUrl = lib.IfThenElse(agencyFareUrl != "", &agencyFareUrl, nil)
	agency.AgencyLang = lib.IfThenElse(agencyLang != "", &agencyLang, nil)
	agency.AgencyPhone = lib.IfThenElse(agencyPhone != "", &agencyPhone, nil)
	agency.AgencyId = lib.IfThenElse(agencyId != "", &agencyId, nil)
	
	agency.AgencyName = agencyName
	agency.AgencyTimezone = agencyTimezone
	agency.AgencyUrl = agencyUrl

	return agency
}