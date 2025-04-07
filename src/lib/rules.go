package lib

import (
	"main/src/models"
)

func CreateAgencyFromMap(m map[string]string) (a models.Agency, errors []string) {
	
	errors = []string{}
	item := models.Agency{}

	//Convert Optional Values
	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone string

	ParseStringToPrimitive(m["agency_email"], &agencyEmail, &errors)
	ParseStringToPrimitive(m["agency_fare_url"], &agencyFareUrl, &errors)
	ParseStringToPrimitive(m["agency_lang"], &agencyLang, &errors)
	ParseStringToPrimitive(m["agency_phone"], &agencyPhone, &errors)

	item.AgencyEmail = &agencyEmail
	item.AgencyFareUrl = &agencyFareUrl
	item.AgencyLang = &agencyLang
	item.AgencyPhone = &agencyPhone
	
	//Convert Required Values
	ParseStringToPrimitive(m["agency_timezone"], &item.AgencyTimezone, &errors)
	ParseStringToPrimitive(m["agency_name"], &item.AgencyName, &errors)
	ParseStringToPrimitive(m["agency_id"], &item.AgencyId, &errors)
	ParseStringToPrimitive(m["agency_url"], &item.AgencyUrl, &errors)

	return item, errors
}
func CreateStopFromMap(m map[string]string) (s models.Stop, errors []string) {
	
	errors = []string{}
	item := models.Stop{}

	//Convert Optional Primitive Values
	var levelId, parentStation, platformCode, stopCode, stopDesc, stopName, stopTimezone, stopUrl, wheelchairBoarding, zoneId string
	var locationType, wheelchairBoardingType uint8

	ParseStringToPrimitive(m["level_id"], &levelId, &errors)
	ParseStringToPrimitive(m["parent_station"], &parentStation, &errors)
	ParseStringToPrimitive(m["platform_code"], &platformCode, &errors)
	ParseStringToPrimitive(m["stop_code"], &stopCode, &errors)
	ParseStringToPrimitive(m["stop_desc"], &stopDesc, &errors)
	ParseStringToPrimitive(m["stop_name"], &stopName, &errors)
	ParseStringToPrimitive(m["stop_timezone"], &stopTimezone, &errors)
	ParseStringToPrimitive(m["stop_url"], &stopUrl, &errors)
	ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoarding, &errors)
	ParseStringToPrimitive(m["zone_id"], &zoneId, &errors)
	ParseStringToPrimitive(m["location_type"], &locationType, &errors)
	ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoardingType, &errors)

	item.LevelId = &levelId
	item.ParentStation = &parentStation
	item.PlatformCode = &platformCode
	item.StopCode = &stopCode
	item.StopDesc = &stopDesc
	item.StopName = &stopName
	item.StopTimezone = &stopTimezone
	item.StopUrl = &stopUrl
	item.ZoneId = &zoneId
	item.LocationType = &locationType
	item.WheelchairBoarding = &wheelchairBoardingType

	//Convert Required Values
	ParseStringToPrimitive(m["stop_id"], &item.StopId, &errors)
	ParseStringToPrimitive(m["stop_lat"], &item.StopLat, &errors)
	ParseStringToPrimitive(m["stop_lon"], &item.StopLon, &errors)

	return item, errors
}