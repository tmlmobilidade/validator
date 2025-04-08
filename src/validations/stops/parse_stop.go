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
	stopIds := make(map[string]bool)

	for i, stop := range gtfsData["stop"] {
		stopMessages := parseStop(stop)

		// Check for duplicate stop IDs
		if stopId, exists := stop["stop_id"]; exists && stopId != "" {
			if stopIds[stopId] {
				messages = append(messages, types.Message{
					Field:        "stop_id",
					FileName:     "stop.txt",
					Message:      "Duplicate stop_id found. Stop IDs must be unique.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			stopIds[stopId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range stopMessages {
			msg.Row = i + 1
			msg.FileName = "stop.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return messages
}

func parseStop(m map[string]string) []types.Message {
	var messages []types.Message
	var parsingErrors []string
	item := types.Stop{}

	//Convert Optional Primitive Values
	var levelId, parentStation, platformCode, stopCode, stopDesc, stopName, stopTimezone, stopUrl, wheelchairBoarding, zoneId string
	var locationType, wheelchairBoardingType uint8
	var stopLat, stopLon float32

	lib.ParseStringToPrimitive(m["level_id"], &levelId, &parsingErrors)
	lib.ParseStringToPrimitive(m["parent_station"], &parentStation, &parsingErrors)
	lib.ParseStringToPrimitive(m["platform_code"], &platformCode, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_code"], &stopCode, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_desc"], &stopDesc, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_name"], &stopName, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_timezone"], &stopTimezone, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_url"], &stopUrl, &parsingErrors)
	lib.ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoarding, &parsingErrors)
	lib.ParseStringToPrimitive(m["zone_id"], &zoneId, &parsingErrors)
	lib.ParseStringToPrimitive(m["location_type"], &locationType, &parsingErrors)
	lib.ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoardingType, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_lat"], &stopLat, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_lon"], &stopLon, &parsingErrors)

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
	lib.ParseStringToPrimitive(m["stop_id"], &item.StopId, &parsingErrors)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A", //TODO: Add field name
				Message: err,
			})
		}
	}

	// Validate Values

	// Validate Required stop_id
	if item.StopId == "" {
		messages = append(messages, types.Message{
			Field:   "stop_id",
			Message: "Stop ID is required and must be unique.",
		})
	}

	// Validate stop_name based on location_type
	if *item.LocationType == 0 || *item.LocationType == 1 || *item.LocationType == 2 {
		if *item.StopName == "" {
			messages = append(messages, types.Message{
				Field:   "stop_name",
				Message: "Stop name is required for stops (location_type=0), stations (location_type=1), and entrances/exits (location_type=2).",
			})
		}
	}

	// Validate stop_lat and stop_lon based on location_type
	if *item.LocationType == 0 || *item.LocationType == 1 || *item.LocationType == 2 {
		if item.StopLat == nil || *item.StopLat == 0 {
			messages = append(messages, types.Message{
				Field:   "stop_lat",
				Message: "Stop latitude is required for stops (location_type=0), stations (location_type=1), and entrances/exits (location_type=2).",
			})
		}
		if item.StopLon == nil || *item.StopLon == 0 {
			messages = append(messages, types.Message{
				Field:   "stop_lon",
				Message: "Stop longitude is required for stops (location_type=0), stations (location_type=1), and entrances/exits (location_type=2).",
			})
		}
	}

	// Validate parent_station based on location_type
	if *item.LocationType == 2 || *item.LocationType == 3 || *item.LocationType == 4 {
		if item.ParentStation == nil || *item.ParentStation == "" {
			messages = append(messages, types.Message{
				Field:   "parent_station",
				Message: "Parent station is required for entrances (location_type=2), generic nodes (location_type=3), and boarding areas (location_type=4).",
			})
		}
	} else if *item.LocationType == 1 && item.ParentStation != nil && *item.ParentStation != "" {
		messages = append(messages, types.Message{
			Field:   "parent_station",
			Message: "Parent station must be empty for stations (location_type=1).",
		})
	}

	// Validate location_type enum values
	if item.LocationType != nil && (*item.LocationType < 0 || *item.LocationType > 4) {
		messages = append(messages, types.Message{
			Field:   "location_type",
			Message: "Invalid location_type. Valid values are 0-4.",
		})
	}

	// Validate wheelchair_boarding enum values
	if item.WheelchairBoarding != nil && (*item.WheelchairBoarding < 0 || *item.WheelchairBoarding > 2) {
		messages = append(messages, types.Message{
			Field:   "wheelchair_boarding",
			Message: "Invalid wheelchair_boarding value. Valid values are 0-2.",
		})
	}

	// Validate URLs if provided
	if item.StopUrl != nil && *item.StopUrl != "" {
		if urlErrors := lib.ValidateUrl(*item.StopUrl); urlErrors != "" {
			messages = append(messages, types.Message{
				Field:   "stop_url",
				Message: urlErrors,
			})
		}
	}

	return messages
}
