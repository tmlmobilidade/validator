package stops

import (
	"main/src/lib"
	"main/src/types"
)

type parseStopValidation struct {
	*types.Validation
}

func NewParseStopValidation(severity *types.Severity) *parseStopValidation {

	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseStopValidation{
		Validation: &types.Validation{
			ID:          "parse_stop",
			Description: "Validate stop data",
			Severity:    s,
		},
	}
}

func (v *parseStopValidation) Validate(gtfsData types.Gtfs) []types.Message {
	var messages []types.Message

	for i, stop := range gtfsData["stop"] {
		_, errs := parseStop(stop)
		for _, err := range errs {
			messages = append(messages, types.Message{
				Field:        "N/A",
				FileName:     "stop.txt",
				Message:      err,
				Row:          i,
				Severity:     v.Severity,
				ValidationID: v.ID,
			})
		}
	}
	return messages
}

func parseStop(m map[string]string) (s types.Stop, errors []string) {
	errors = []string{}
	item := types.Stop{}

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

	return item, errors
}
