package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseAgency(rawAgency map[string]string, row int, gtfs *types.Gtfs) types.Agency {
	var (
		agency                                           types.Agency = types.Agency{}
		agencyName, agencyUrl, agencyTimezone, agencyId, agencyLang, agencyPhone, agencyFareUrl, agencyEmail                     string
		messages                                       []types.Message
	)

	stringFields := map[string]*string{
		"agency_name":         &agencyName,
		"agency_url":        &agencyUrl,
		"agency_timezone":      &agencyTimezone,
		"agency_id":   &agencyId,
		"agency_lang":   &agencyLang,
		"agency_phone": &agencyPhone,
		"agency_fare_url": &agencyFareUrl,
		"agency_email": &agencyEmail,
	}


	// Helper to collect error messages
	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "agency.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "agency_parse",
		})
	}

	// Parse string fields
	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawAgency[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	// If there are any errors, return an empty trip
	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return agency
	}

	// Required fields
	agency.AgencyName = agencyName
	agency.AgencyUrl = agencyUrl
	agency.AgencyTimezone = agencyTimezone

	agency.AgencyId = lib.IfThenElse(rawAgency["agency_id"] != "", &agencyId, nil)
	agency.AgencyLang = lib.IfThenElse(rawAgency["agency_lang"] != "", &agencyLang, nil)
	agency.AgencyPhone = lib.IfThenElse(rawAgency["agency_phone"] != "", &agencyPhone, nil)
	agency.AgencyFareUrl = lib.IfThenElse(rawAgency["agency_fare_url"] != "", &agencyFareUrl, nil)
	agency.AgencyEmail = lib.IfThenElse(rawAgency["agency_email"] != "", &agencyEmail, nil)

	return agency
}