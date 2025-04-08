package agency

import (
	"main/src/lib"
	"main/src/types"
)

type ParseAgencyValidation struct {
	*types.Validation
}

func NewParseAgencyValidation(severity *types.Severity) *ParseAgencyValidation {

	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &ParseAgencyValidation{
		Validation: &types.Validation{
			ID:          "parse_agency",
			Description: "Validate agency data",
			Severity:    s,
		},
	}
}

func (v *ParseAgencyValidation) Validate(gtfsData types.Gtfs) []types.Message {
	var messages []types.Message
	agencyIds := make(map[string]bool)

	for i, agency := range gtfsData["agency"] {
		agencyMessages := parseAgency(agency, len(gtfsData["agency"]))

		// Check for duplicate agency IDs
		if agencyId, exists := agency["agency_id"]; exists && agencyId != "" {
			if agencyIds[agencyId] {
				messages = append(messages, types.Message{
					Field:        "agency_id",
					FileName:     "agency.txt",
					Message:      "Duplicate agency_id found. Agency IDs must be unique.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			agencyIds[agencyId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range agencyMessages {
			msg.Row = i + 1
			msg.FileName = "agency.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return messages
}

func parseAgency(m map[string]string, totalAgencies int) []types.Message {
	var messages []types.Message
	item := types.Agency{}

	//Convert Optional Values
	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone, agencyId string

	lib.ParseStringToPrimitive(m["agency_email"], &agencyEmail, nil)
	lib.ParseStringToPrimitive(m["agency_fare_url"], &agencyFareUrl, nil)
	lib.ParseStringToPrimitive(m["agency_lang"], &agencyLang, nil)
	lib.ParseStringToPrimitive(m["agency_phone"], &agencyPhone, nil)
	lib.ParseStringToPrimitive(m["agency_id"], &agencyId, nil)

	item.AgencyEmail = &agencyEmail
	item.AgencyFareUrl = &agencyFareUrl
	item.AgencyLang = &agencyLang
	item.AgencyPhone = &agencyPhone
	item.AgencyId = &agencyId

	//Convert Required Values
	lib.ParseStringToPrimitive(m["agency_timezone"], &item.AgencyTimezone, nil)
	lib.ParseStringToPrimitive(m["agency_name"], &item.AgencyName, nil)
	lib.ParseStringToPrimitive(m["agency_url"], &item.AgencyUrl, nil)

	// Validate Values
	if item.AgencyTimezone == "" {
		messages = append(messages, types.Message{
			Field:   "agency_timezone",
			Message: "Agency timezone is required.",
		})
	} else if tzErrors := lib.ValidateTimezone(item.AgencyTimezone); len(tzErrors) > 0 {
		for _, err := range tzErrors {
			messages = append(messages, types.Message{
				Field:   "agency_timezone",
				Message: err,
			})
		}
	}

	// Validate Agency URL
	if item.AgencyUrl == "" {
		messages = append(messages, types.Message{
			Field:   "agency_url",
			Message: "Agency URL is required.",
		})
	} else if urlErrors := lib.ValidateUrl(item.AgencyUrl); len(urlErrors) > 0 {
		for _, err := range urlErrors {
			messages = append(messages, types.Message{
				Field:   "agency_url",
				Message: err,
			})
		}
	}

	// Validate Agency Name
	if item.AgencyName == "" {
		messages = append(messages, types.Message{
			Field:   "agency_name",
			Message: "Agency name is required.",
		})
	}

	// Validate Agency ID
	if totalAgencies > 1 && *item.AgencyId == "" {
		messages = append(messages, types.Message{
			Field:   "agency_id",
			Message: "Agency ID is required when the dataset contains data for multiple transit agencies.",
		})
	}

	// Validate Agency Phone
	if item.AgencyPhone != nil && *item.AgencyPhone != "" {
		if phoneErrors := lib.ValidatePhone(*item.AgencyPhone); len(phoneErrors) > 0 {
			for _, err := range phoneErrors {
				messages = append(messages, types.Message{
					Field:   "agency_phone",
					Message: err,
				})
			}
		}
	}

	// Validate Agency Email
	if item.AgencyEmail != nil && *item.AgencyEmail != "" {
		if emailErrors := lib.ValidateEmail(*item.AgencyEmail); len(emailErrors) > 0 {
			for _, err := range emailErrors {
				messages = append(messages, types.Message{
					Field:   "agency_email",
					Message: err,
				})
			}
		}
	}

	// Validate Agency Fare URL
	if item.AgencyFareUrl != nil && *item.AgencyFareUrl != "" {
		if urlErrors := lib.ValidateUrl(*item.AgencyFareUrl); len(urlErrors) > 0 {
			for _, err := range urlErrors {
				messages = append(messages, types.Message{
					Field:   "agency_fare_url",
					Message: err,
				})
			}
		}
	}

	// Validate Agency Language
	if item.AgencyLang != nil && *item.AgencyLang != "" {
		if langErrors := lib.ValidateLanguage(*item.AgencyLang); len(langErrors) > 0 {
			for _, err := range langErrors {
				messages = append(messages, types.Message{
					Field:   "agency_lang",
					Message: err,
				})
			}
		}
	}

	return messages
}
