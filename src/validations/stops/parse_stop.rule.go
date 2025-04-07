package stops

import (
	"main/src/lib"
	"main/src/models"
)

func ParseStop(m map[string]string) (s models.Stop, errors []string) {
	errors = []string{}
	item := models.Stop{}

	//Convert Optional Primitive Values
	var levelId, parentStation, platformCode, stopCode, stopDesc, stopName, stopTimezone, stopUrl, wheelchairBoarding, zoneId string
	var locationType, wheelchairBoardingType uint8
	var stopLat, stopLon float32

	lib.ParseStringToPrimitive(m["level_id"], &levelId, &errors)
	lib.ParseStringToPrimitive(m["parent_station"], &parentStation, &errors)
	lib.ParseStringToPrimitive(m["platform_code"], &platformCode, &errors)
	lib.ParseStringToPrimitive(m["stop_code"], &stopCode, &errors)
	lib.ParseStringToPrimitive(m["stop_desc"], &stopDesc, &errors)
	lib.ParseStringToPrimitive(m["stop_name"], &stopName, &errors)
	lib.ParseStringToPrimitive(m["stop_timezone"], &stopTimezone, &errors)
	lib.ParseStringToPrimitive(m["stop_url"], &stopUrl, &errors)
	lib.ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoarding, &errors)
	lib.ParseStringToPrimitive(m["zone_id"], &zoneId, &errors)
	lib.ParseStringToPrimitive(m["location_type"], &locationType, &errors)
	lib.ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoardingType, &errors)
	lib.ParseStringToPrimitive(m["stop_lat"], &stopLat, &errors)
	lib.ParseStringToPrimitive(m["stop_lon"], &stopLon, &errors)

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
	item.StopLat = &stopLat

	//Convert Required Values
	lib.ParseStringToPrimitive(m["stop_id"], &item.StopId, &errors)
	lib.ParseStringToPrimitive(m["stop_lat"], &item.StopLat, &errors) // Will panic if not float32

	return item, errors
}
