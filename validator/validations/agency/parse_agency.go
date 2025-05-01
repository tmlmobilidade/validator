package agency

import (
	"main/lib"
	"main/types"
)

func ParseAgencyValidation(severity *types.Severity, rawAgency map[string]string, row int, gtfs *types.Gtfs) (agency types.Agency, messages []types.Message) {

	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	var parsingErrors []string

	//Convert Optional Values
	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone, agencyId string

	lib.ParseStringToPrimitive(rawAgency["agency_email"], &agencyEmail, &parsingErrors)
	lib.ParseStringToPrimitive(rawAgency["agency_fare_url"], &agencyFareUrl, &parsingErrors)
	lib.ParseStringToPrimitive(rawAgency["agency_lang"], &agencyLang, &parsingErrors)
	lib.ParseStringToPrimitive(rawAgency["agency_phone"], &agencyPhone, &parsingErrors)
	lib.ParseStringToPrimitive(rawAgency["agency_id"], &agencyId, &parsingErrors)

	agency.AgencyEmail = lib.IfThenElse(agencyEmail != "", &agencyEmail, nil)
	agency.AgencyFareUrl = lib.IfThenElse(agencyFareUrl != "", &agencyFareUrl, nil)
	agency.AgencyLang = lib.IfThenElse(agencyLang != "", &agencyLang, nil)
	agency.AgencyPhone = lib.IfThenElse(agencyPhone != "", &agencyPhone, nil)
	agency.AgencyId = lib.IfThenElse(agencyId != "", &agencyId, nil)

	//Convert Required Values
	lib.ParseStringToPrimitive(rawAgency["agency_timezone"], &agency.AgencyTimezone, &parsingErrors)
	lib.ParseStringToPrimitive(rawAgency["agency_name"], &agency.AgencyName, &parsingErrors)
	lib.ParseStringToPrimitive(rawAgency["agency_url"], &agency.AgencyUrl, &parsingErrors)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A", //TODO: Add field name
				Message: err,
				FileName: "agency.txt",
				Rows: []int{row},
				Severity: s,
				ValidationID: "agency_parse",
			})
		}
	}

	return agency, messages
}

// func (v *ParseAgencyValidation) Validate(gtfs types.Gtfs) (agencies []types.Agency, messages []types.Message) {
// 	agencyIds := make(map[string]bool)

// 	for i, agency := range gtfs.Files["agency"] {
// 		agency, agencyMessages := parseAgency(agency, len(gtfs.Files["agency"]))
// 		agencies = append(agencies, agency)

// 		// Check for duplicate agency IDs
// 		if agency.AgencyId != nil && *agency.AgencyId != "" {
// 			if agencyIds[*agency.AgencyId] {
// 				messages = append(messages, types.Message{
// 					Field:        "agency_id",
// 					FileName:     "agency.txt",
// 					Message:      "Duplicate agency_id found. Agency IDs must be unique.",
// 					Row:          i + 1,
// 					Severity:     v.Severity,
// 					ValidationID: v.ID,
// 				})
// 			}
// 			agencyIds[*agency.AgencyId] = true
// 		}

// 		// Update row number and other fields for each message
// 		for _, msg := range agencyMessages {
// 			msg.Row = i + 1
// 			msg.FileName = "agency.txt"
// 			msg.Severity = v.Severity
// 			msg.ValidationID = v.ID
// 			messages = append(messages, msg)
// 		}
// 	}
// 	return agencies, messages
// }

// func parseAgency(m map[string]string, totalAgencies int) (agency types.Agency, messages []types.Message) {
// 	var parsingErrors []string

// 	//Convert Optional Values
// 	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone, agencyId string

// 	lib.ParseStringToPrimitive(m["agency_email"], &agencyEmail, &parsingErrors)
// 	lib.ParseStringToPrimitive(m["agency_fare_url"], &agencyFareUrl, &parsingErrors)
// 	lib.ParseStringToPrimitive(m["agency_lang"], &agencyLang, &parsingErrors)
// 	lib.ParseStringToPrimitive(m["agency_phone"], &agencyPhone, &parsingErrors)
// 	lib.ParseStringToPrimitive(m["agency_id"], &agencyId, &parsingErrors)

// 	agency.AgencyEmail = lib.IfThenElse(agencyEmail != "", &agencyEmail, nil)
// 	agency.AgencyFareUrl = lib.IfThenElse(agencyFareUrl != "", &agencyFareUrl, nil)
// 	agency.AgencyLang = lib.IfThenElse(agencyLang != "", &agencyLang, nil)
// 	agency.AgencyPhone = lib.IfThenElse(agencyPhone != "", &agencyPhone, nil)
// 	agency.AgencyId = lib.IfThenElse(agencyId != "", &agencyId, nil)

// 	//Convert Required Values
// 	lib.ParseStringToPrimitive(m["agency_timezone"], &agency.AgencyTimezone, &parsingErrors)
// 	lib.ParseStringToPrimitive(m["agency_name"], &agency.AgencyName, &parsingErrors)
// 	lib.ParseStringToPrimitive(m["agency_url"], &agency.AgencyUrl, &parsingErrors)

// 	if len(parsingErrors) > 0 {
// 		for _, err := range parsingErrors {
// 			messages = append(messages, types.Message{
// 				Field:   "N/A", //TODO: Add field name
// 				Message: err,
// 			})
// 		}
// 	}
// 	// Validate Values
// 	if agency.AgencyTimezone == "" {
// 		messages = append(messages, types.Message{
// 			Field:   "agency_timezone",
// 			Message: "Agency timezone is required.",
// 		})
// 	} else if tzErrors := lib.ValidateTimezone(agency.AgencyTimezone); tzErrors != "" {
// 		messages = append(messages, types.Message{
// 			Field:   "agency_timezone",
// 			Message: tzErrors,
// 		})
// 	}

// 	// Validate Agency URL
// 	if agency.AgencyUrl == "" {
// 		messages = append(messages, types.Message{
// 			Field:   "agency_url",
// 			Message: "Agency URL is required.",
// 		})
// 	} else if urlErrors := lib.ValidateUrl(agency.AgencyUrl); urlErrors != "" {
// 		messages = append(messages, types.Message{
// 			Field:   "agency_url",
// 			Message: urlErrors,
// 		})
// 	}

// 	// Validate Agency Name
// 	if agency.AgencyName == "" {
// 		messages = append(messages, types.Message{
// 			Field:   "agency_name",
// 			Message: "Agency name is required.",
// 		})
// 	}

// 	// Validate Agency ID
// 	if totalAgencies > 1 && agency.AgencyId == nil {
// 		messages = append(messages, types.Message{
// 			Field:   "agency_id",
// 			Message: "Agency ID is required when the dataset contains data for multiple transit agencies.",
// 		})
// 	}

// 	// Validate Agency Phone
// 	if agency.AgencyPhone != nil && *agency.AgencyPhone != "" {
// 		if phoneErrors := lib.ValidatePhone(*agency.AgencyPhone); phoneErrors != "" {
// 			messages = append(messages, types.Message{
// 				Field:   "agency_phone",
// 				Message: phoneErrors,
// 			})
// 		}
// 	}

// 	// Validate Agency Email
// 	if agency.AgencyEmail != nil && *agency.AgencyEmail != "" {
// 		if emailErrors := lib.ValidateEmail(*agency.AgencyEmail); emailErrors != "" {
// 			messages = append(messages, types.Message{
// 				Field:   "agency_email",
// 				Message: emailErrors,
// 			})
// 		}
// 	}

// 	// Validate Agency Fare URL
// 	if agency.AgencyFareUrl != nil && *agency.AgencyFareUrl != "" {
// 		if urlErrors := lib.ValidateUrl(*agency.AgencyFareUrl); urlErrors != "" {
// 			messages = append(messages, types.Message{
// 				Field:   "agency_fare_url",
// 				Message: urlErrors,
// 			})
// 		}
// 	}

// 	// Validate Agency Language
// 	if agency.AgencyLang != nil && *agency.AgencyLang != "" {
// 		if langErrors := lib.ValidateLanguage(*agency.AgencyLang); langErrors != "" {
// 			messages = append(messages, types.Message{
// 				Field:   "agency_lang",
// 				Message: langErrors,
// 			})
// 		}
// 	}

// 	return agency, messages
// }
