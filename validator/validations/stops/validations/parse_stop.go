package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

// ParseStop parses a row from stops.txt into a Stop struct, following gtfs-parser-validation best practices
func ParseStop(rawStop map[string]string, row int) types.Stop {
	var (
		stop types.Stop = types.Stop{}
		stopId string
		stopCode, stopName, stopDesc, zoneId, stopUrl, parentStation, stopTimezone, levelId, platformCode string
		locationType, wheelchairBoarding int
		stopLat, stopLon float32
		messages []types.Message
	)

	stringFields := map[string]*string{
		"stop_id": &stopId,
		"stop_code": &stopCode,
		"stop_name": &stopName,
		"stop_desc": &stopDesc,
		"zone_id": &zoneId,
		"stop_url": &stopUrl,
		"parent_station": &parentStation,
		"stop_timezone": &stopTimezone,
		"level_id": &levelId,
		"platform_code": &platformCode,
	}

	intFields := map[string]*int{
		"location_type": &locationType,
		"wheelchair_boarding": &wheelchairBoarding,
	}

	float32Fields := map[string]*float32{
		"stop_lat": &stopLat,
		"stop_lon": &stopLon,
	}

	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "stops_parse",
		})
	}

	// Parse string fields
	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawStop[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}
	// Parse int fields
	for field, target := range intFields {
		if errMsg := lib.ParseStringToPrimitive(rawStop[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}
	// Parse float32 fields
	for field, target := range float32Fields {
		if errMsg := lib.ParseStringToPrimitive(rawStop[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return stop
	}

	// Assign fields
	stop.StopId = lib.IfThenElse(rawStop["stop_id"] != "", &stopId, nil)
	stop.StopCode = lib.IfThenElse(rawStop["stop_code"] != "", &stopCode, nil)
	stop.StopName = lib.IfThenElse(rawStop["stop_name"] != "", &stopName, nil)
	stop.StopDesc = lib.IfThenElse(rawStop["stop_desc"] != "", &stopDesc, nil)
	stop.ZoneId = lib.IfThenElse(rawStop["zone_id"] != "", &zoneId, nil)
	stop.StopUrl = lib.IfThenElse(rawStop["stop_url"] != "", &stopUrl, nil)
	stop.ParentStation = lib.IfThenElse(rawStop["parent_station"] != "", &parentStation, nil)
	stop.StopTimezone = lib.IfThenElse(rawStop["stop_timezone"] != "", &stopTimezone, nil)
	stop.LevelId = lib.IfThenElse(rawStop["level_id"] != "", &levelId, nil)
	stop.PlatformCode = lib.IfThenElse(rawStop["platform_code"] != "", &platformCode, nil)
	stop.LocationType = lib.IfThenElse(rawStop["location_type"] != "", &locationType, nil)
	stop.WheelchairBoarding = lib.IfThenElse(rawStop["wheelchair_boarding"] != "", &wheelchairBoarding, nil)
	stop.StopLat = lib.IfThenElse(rawStop["stop_lat"] != "", &stopLat, nil)
	stop.StopLon = lib.IfThenElse(rawStop["stop_lon"] != "", &stopLon, nil)

	return stop
}