package trips

import (
	"main/src/lib"
	"main/src/types"
)

type parseTripValidation struct {
	*types.Validation
}

func NewParseTripValidation(severity *types.Severity) *parseTripValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseTripValidation{
		Validation: &types.Validation{
			ID:          "parse_trip",
			Description: "Validate trip data",
			Severity:    s,
		},
	}
}

func (v *parseTripValidation) Validate(gtfs types.Gtfs) (trips []types.Trip, messages []types.Message) {
	tripIds := make(map[string]bool)

	// Check if any routes have continuous pickup/dropoff behavior
	hasContinuousPickupDropoff := false
	for _, route := range gtfs.Files["routes"] {
		if (route["continuous_pickup"] != "" && route["continuous_pickup"] != "0") ||
			(route["continuous_drop_off"] != "" && route["continuous_drop_off"] != "0") {
			hasContinuousPickupDropoff = true
			break
		}
	}

	// Check if any stop_times have continuous pickup/dropoff behavior
	hasStopTimesContinuousPickupDropoff := false
	for _, stopTime := range gtfs.Files["stop_times"] {
		if (stopTime["continuous_pickup"] != "" && stopTime["continuous_pickup"] != "0") ||
			(stopTime["continuous_drop_off"] != "" && stopTime["continuous_drop_off"] != "0") {
			hasStopTimesContinuousPickupDropoff = true
			break
		}
	}

	for i, trip := range gtfs.Files["trips"] {
		trip, tripMessages := parseTrip(trip, gtfs.IdMap["routes"], gtfs.IdMap["calendar"], gtfs.IdMap["shapes"], hasContinuousPickupDropoff, hasStopTimesContinuousPickupDropoff)
		trips = append(trips, trip)

		// Check for duplicate trip IDs
		if trip.TripId != "" {
			if tripIds[trip.TripId] {
				messages = append(messages, types.Message{
					Field:        "trip_id",
					FileName:     "trips.txt",
					Message:      "Duplicate trip_id found. Trip IDs must be unique.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			tripIds[trip.TripId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range tripMessages {
			msg.Row = i + 1
			msg.FileName = "trips.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return trips, messages
}

func parseTrip(m map[string]string, routeIds map[string]int, serviceIds map[string]int, shapeIds map[string]int, hasContinuousPickupDropoff bool, hasStopTimesContinuousPickupDropoff bool) (trip types.Trip, messages []types.Message) {
	var parsingErrors []string

	// Convert Optional Primitive Values
	var bikesAllowed int
	var directionId bool
	var wheelchairAccessible string
	var blockId, shapeId, tripHeadsign, tripShortName string

	lib.ParseStringToPrimitive(m["bikes_allowed"], &bikesAllowed, &parsingErrors)
	lib.ParseStringToPrimitive(m["direction_id"], &directionId, &parsingErrors)
	lib.ParseStringToPrimitive(m["wheelchair_accessible"], &wheelchairAccessible, &parsingErrors)
	lib.ParseStringToPrimitive(m["block_id"], &blockId, &parsingErrors)
	lib.ParseStringToPrimitive(m["shape_id"], &shapeId, &parsingErrors)
	lib.ParseStringToPrimitive(m["trip_headsign"], &tripHeadsign, &parsingErrors)
	lib.ParseStringToPrimitive(m["trip_short_name"], &tripShortName, &parsingErrors)

	trip.BikesAllowed = lib.IfThenElse(m["bikes_allowed"] != "", &bikesAllowed, nil)
	trip.DirectionId = lib.IfThenElse(m["direction_id"] != "", &directionId, nil)
	trip.WheelchairAccessible = lib.IfThenElse(m["wheelchair_accessible"] != "", &wheelchairAccessible, nil)
	trip.BlockId = lib.IfThenElse(m["block_id"] != "", &blockId, nil)
	trip.ShapeId = lib.IfThenElse(m["shape_id"] != "", &shapeId, nil)
	trip.TripHeadsign = lib.IfThenElse(m["trip_headsign"] != "", &tripHeadsign, nil)
	trip.TripShortName = lib.IfThenElse(m["trip_short_name"] != "", &tripShortName, nil)

	// Convert Required Values
	lib.ParseStringToPrimitive(m["trip_id"], &trip.TripId, &parsingErrors)
	lib.ParseStringToPrimitive(m["route_id"], &trip.RouteId, &parsingErrors)
	lib.ParseStringToPrimitive(m["service_id"], &trip.ServiceId, &parsingErrors)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A", //TODO: Add field name
				Message: err,
			})
		}
	}

	// Validate Values
	// Validate Required trip_id
	if trip.TripId == "" {
		messages = append(messages, types.Message{
			Field:   "trip_id",
			Message: "Trip ID is required and must be unique.",
		})
	}

	// Validate Required route_id
	if trip.RouteId == "" {
		messages = append(messages, types.Message{
			Field:   "route_id",
			Message: "Route ID is required.",
		})
	} else {
		_, ok := routeIds[trip.RouteId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "route_id",
				Message: "Route ID must reference a valid route_id from routes.txt.",
			})
		}
	}

	// Validate Required service_id
	if trip.ServiceId == "" {
		messages = append(messages, types.Message{
			Field:   "service_id",
			Message: "Service ID is required.",
		})
	} else {
		_, ok := serviceIds[trip.ServiceId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "service_id",
				Message: "Service ID must reference a valid service_id from calendar.txt or calendar_dates.txt.",
			})
		}
	}

	// Validate shape_id if provided
	if trip.ShapeId != nil && *trip.ShapeId != "" {
		_, ok := shapeIds[*trip.ShapeId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "shape_id",
				Message: "Shape ID must reference a valid shape_id from shapes.txt.",
			})
		}
	}

	// Validate shape_id is required if continuous pickup/dropoff behavior is defined
	if (hasContinuousPickupDropoff || hasStopTimesContinuousPickupDropoff) && (trip.ShapeId == nil || *trip.ShapeId == "") {
		messages = append(messages, types.Message{
			Field:   "shape_id",
			Message: "Shape ID is required when continuous pickup or drop-off behavior is defined in routes.txt or stop_times.txt.",
		})
	}

	// Validate direction_id enum values
	if trip.DirectionId != nil {
		// direction_id is a boolean in the struct, so we don't need to validate enum values
		// The value will be true for 1 and false for 0
	}

	// Validate bikes_allowed enum values
	if trip.BikesAllowed != nil {
		validBikesAllowed := map[int]bool{0: true, 1: true, 2: true}
		if !validBikesAllowed[*trip.BikesAllowed] {
			messages = append(messages, types.Message{
				Field:   "bikes_allowed",
				Message: "Invalid bikes_allowed value. Valid values are 0, 1, 2.",
			})
		}
	}

	// Validate wheelchair_accessible enum values
	if trip.WheelchairAccessible != nil && *trip.WheelchairAccessible != "" {
		validWheelchairAccessible := map[string]bool{"0": true, "1": true, "2": true}
		if !validWheelchairAccessible[*trip.WheelchairAccessible] {
			messages = append(messages, types.Message{
				Field:   "wheelchair_accessible",
				Message: "Invalid wheelchair_accessible value. Valid values are 0, 1, 2.",
			})
		}
	}

	return trip, messages
}
