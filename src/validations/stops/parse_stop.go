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

func (v *parseStopValidation) Validate(gtfsData types.Gtfs) (stops []types.Stop, messages []types.Message) {
	stopIds := make(map[string]bool)

	for i, stop := range gtfsData["stops"] {
		stop, stopMessages := parseStop(stop)
		stops = append(stops, stop)

		// Check for duplicate stop IDs
		if stop.StopId != "" {
			if stopIds[stop.StopId] {
				messages = append(messages, types.Message{
					Field:        "stop_id",
					FileName:     "stop.txt",
					Message:      "Duplicate stop_id found. Stop IDs must be unique.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			stopIds[stop.StopId] = true
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
	return stops, messages
}

func parseStop(m map[string]string) (stop types.Stop, messages []types.Message) {
	var parsingErrors []string

	//Convert Optional Primitive Values
	var levelId, parentStation, platformCode, stopCode, stopDesc, stopName, stopTimezone, stopUrl, wheelchairBoarding, zoneId string
	var wheelchairBoardingType, locationType int
	var stopLat, stopLon float32

	lib.ParseStringToPrimitive(m["location_type"], &locationType, &parsingErrors)
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
	lib.ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoardingType, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_lat"], &stopLat, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_lon"], &stopLon, &parsingErrors)

	stop.LevelId = lib.IfThenElse(m["level_id"] != "", &levelId, nil)
	stop.ParentStation = lib.IfThenElse(m["parent_station"] != "", &parentStation, nil)
	stop.PlatformCode = lib.IfThenElse(m["platform_code"] != "", &platformCode, nil)
	stop.StopCode = lib.IfThenElse(m["stop_code"] != "", &stopCode, nil)
	stop.StopDesc = lib.IfThenElse(m["stop_desc"] != "", &stopDesc, nil)
	stop.StopName = lib.IfThenElse(m["stop_name"] != "", &stopName, nil)
	stop.StopTimezone = lib.IfThenElse(m["stop_timezone"] != "", &stopTimezone, nil)
	stop.StopUrl = lib.IfThenElse(m["stop_url"] != "", &stopUrl, nil)
	stop.ZoneId = lib.IfThenElse(m["zone_id"] != "", &zoneId, nil)
	stop.StopLat = lib.IfThenElse(m["stop_lat"] != "", &stopLat, nil)
	stop.StopLon = lib.IfThenElse(m["stop_lon"] != "", &stopLon, nil)
	stop.WheelchairBoarding = lib.IfThenElse(m["wheelchair_boarding"] != "", &wheelchairBoardingType, nil)
	stop.LocationType = lib.IfThenElse(m["location_type"] != "", &locationType, nil)

	//Convert Required Values
	lib.ParseStringToPrimitive(m["stop_id"], &stop.StopId, &parsingErrors)

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
	if stop.StopId == "" {
		messages = append(messages, types.Message{
			Field:   "stop_id",
			Message: "Stop ID is required and must be unique.",
		})
	}

	// Validate stop_name based on location_type
	if stop.LocationType != nil && (*stop.LocationType == 0 || *stop.LocationType == 1 || *stop.LocationType == 2) {
		if *stop.StopName == "" {
			messages = append(messages, types.Message{
				Field:   "stop_name",
				Message: "Stop name is required for stops (location_type=0), stations (location_type=1), and entrances/exits (location_type=2).",
			})
		}
	}

	// Validate stop_lat and stop_lon based on location_type
	if stop.LocationType != nil && (*stop.LocationType == 0 || *stop.LocationType == 1 || *stop.LocationType == 2) {
		if stop.StopLat == nil || *stop.StopLat == 0 {
			messages = append(messages, types.Message{
				Field:   "stop_lat",
				Message: "Stop latitude is required for stops (location_type=0), stations (location_type=1), and entrances/exits (location_type=2).",
			})
		}
		if stop.StopLon == nil || *stop.StopLon == 0 {
			messages = append(messages, types.Message{
				Field:   "stop_lon",
				Message: "Stop longitude is required for stops (location_type=0), stations (location_type=1), and entrances/exits (location_type=2).",
			})
		}
	}

	// Validate parent_station based on location_type
	if stop.LocationType != nil && (*stop.LocationType == 2 || *stop.LocationType == 3 || *stop.LocationType == 4) {
		if stop.ParentStation == nil || *stop.ParentStation == "" {
			messages = append(messages, types.Message{
				Field:   "parent_station",
				Message: "Parent station is required for entrances (location_type=2), generic nodes (location_type=3), and boarding areas (location_type=4).",
			})
		}
	} else if stop.LocationType != nil && *stop.LocationType == 1 && stop.ParentStation != nil && *stop.ParentStation != "" {
		messages = append(messages, types.Message{
			Field:   "parent_station",
			Message: "Parent station must be empty for stations (location_type=1).",
		})
	}

	// Validate location_type enum values
	if stop.LocationType != nil && (*stop.LocationType < 0 || *stop.LocationType > 4) {
		messages = append(messages, types.Message{
			Field:   "location_type",
			Message: "Invalid location_type. Valid values are 0-4.",
		})
	}

	// Validate wheelchair_boarding enum values
	if stop.WheelchairBoarding != nil && (*stop.WheelchairBoarding < 0 || *stop.WheelchairBoarding > 2) {
		messages = append(messages, types.Message{
			Field:   "wheelchair_boarding",
			Message: "Invalid wheelchair_boarding value. Valid values are 0-2.",
		})
	}

	// Validate URLs if provided
	if stop.StopUrl != nil && *stop.StopUrl != "" {
		if urlErrors := lib.ValidateUrl(*stop.StopUrl); urlErrors != "" {
			messages = append(messages, types.Message{
				Field:   "stop_url",
				Message: urlErrors,
			})
		}
	}

	return stop, messages
}
