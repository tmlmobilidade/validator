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

	for i, agency := range gtfsData["agency"] {
		_, errs := parseAgency(agency, len(gtfsData["agency"]))
		for _, err := range errs {
			messages = append(messages, types.Message{
				Field:        "N/A",
				FileName:     "agency.txt",
				Message:      err,
				Row:          i + 1,
				Severity:     v.Severity,
				ValidationID: v.ID,
			})
		}
	}
	return messages
}

func parseAgency(m map[string]string, totalAgencies int) (a types.Agency, errors []string) {

	errors = []string{}
	item := types.Agency{}

	//Convert Optional Values
	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone, agencyId string

	lib.ParseStringToPrimitive(m["agency_email"], &agencyEmail, &errors)
	lib.ParseStringToPrimitive(m["agency_fare_url"], &agencyFareUrl, &errors)
	lib.ParseStringToPrimitive(m["agency_lang"], &agencyLang, &errors)
	lib.ParseStringToPrimitive(m["agency_phone"], &agencyPhone, &errors)
	lib.ParseStringToPrimitive(m["agency_id"], &agencyId, &errors)

	item.AgencyEmail = &agencyEmail
	item.AgencyFareUrl = &agencyFareUrl
	item.AgencyLang = &agencyLang
	item.AgencyPhone = &agencyPhone
	item.AgencyId = &agencyId

	//Convert Required Values
	lib.ParseStringToPrimitive(m["agency_timezone"], &item.AgencyTimezone, &errors)
	lib.ParseStringToPrimitive(m["agency_name"], &item.AgencyName, &errors)
	lib.ParseStringToPrimitive(m["agency_url"], &item.AgencyUrl, &errors)

	// Validate Values
	if item.AgencyTimezone == "" {
		errors = append(errors, "Agency timezone is required.")
	} else {
		errors = append(errors, lib.ValidateTimezone(item.AgencyTimezone)...)
	}

	// Validate Agency URL
	if item.AgencyUrl == "" {
		errors = append(errors, "Agency URL is required.")
	} else {
		errors = append(errors, lib.ValidateUrl(item.AgencyUrl)...)
	}

	// Validate Agency Name
	if item.AgencyName == "" {
		errors = append(errors, "Agency name is required.")
	}

	// Validate Agency ID
	if totalAgencies > 1 && *item.AgencyId == "" {
		errors = append(errors, "Agency ID is required when the dataset contains data for multiple transit agencies.")
	}

	// Validate Agency Phone
	if item.AgencyPhone != nil && *item.AgencyPhone != "" {
		errors = append(errors, lib.ValidatePhone(*item.AgencyPhone)...)
	}

	// Validate Agency Email
	if item.AgencyEmail != nil && *item.AgencyEmail != "" {
		errors = append(errors, lib.ValidateEmail(*item.AgencyEmail)...)
	}

	// Validate Agency Fare URL
	if item.AgencyFareUrl != nil && *item.AgencyFareUrl != "" {
		errors = append(errors, lib.ValidateUrl(*item.AgencyFareUrl)...)
	}

	// Validate Agency Language
	if item.AgencyLang != nil && *item.AgencyLang != "" {
		errors = append(errors, lib.ValidateLanguage(*item.AgencyLang)...)
	}

	return item, errors
}
