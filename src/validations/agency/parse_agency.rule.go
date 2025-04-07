package agency

import (
	"main/src/lib"
	"main/src/models"
)

func ParseAgency(m map[string]string) (a models.Agency, errors []string) {

	errors = []string{}
	item := models.Agency{}

	//Convert Optional Values
	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone string

	lib.ParseStringToPrimitive(m["agency_email"], &agencyEmail, &errors)
	lib.ParseStringToPrimitive(m["agency_fare_url"], &agencyFareUrl, &errors)
	lib.ParseStringToPrimitive(m["agency_lang"], &agencyLang, &errors)
	lib.ParseStringToPrimitive(m["agency_phone"], &agencyPhone, &errors)

	item.AgencyEmail = &agencyEmail
	item.AgencyFareUrl = &agencyFareUrl
	item.AgencyLang = &agencyLang
	item.AgencyPhone = &agencyPhone

	//Convert Required Values
	lib.ParseStringToPrimitive(m["agency_timezone"], &item.AgencyTimezone, &errors)
	lib.ParseStringToPrimitive(m["agency_name"], &item.AgencyName, &errors)
	lib.ParseStringToPrimitive(m["agency_id"], &item.AgencyId, &errors)
	lib.ParseStringToPrimitive(m["agency_url"], &item.AgencyUrl, &errors)

	return item, errors
}
