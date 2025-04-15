package stop_times

import (
	"main/validator/lib"
	"main/validator/types"
	"strconv"
)

type parseStopTimeValidation struct {
	*types.Validation
}

func NewParseStopTimeValidation(severity *types.Severity) *parseStopTimeValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseStopTimeValidation{
		Validation: &types.Validation{
			ID:          "parse_stop_time",
			Description: "Validate stop time data",
			Severity:    s,
		},
	}
}

func (v *parseStopTimeValidation) Validate(gtfs types.Gtfs) (stopTimes []types.StopTime, messages []types.Message) {
	// Pre-allocate slices with capacity to avoid resizing
	stopTimes = make([]types.StopTime, 0, len(gtfs.Files["stop_times"]))
	messages = make([]types.Message, 0, len(gtfs.Files["stop_times"])*2) // Estimate 2 messages per stop time

	// Create a map to track unique trip_id + stop_sequence combinations
	uniqueStopSequences := make(map[string]bool, len(gtfs.Files["stop_times"]))

	// Check if any routes have continuous pickup/dropoff behavior
	hasContinuousPickupDropoff := false
	for _, route := range gtfs.Files["routes"] {
		if (route["continuous_pickup"] != "" && route["continuous_pickup"] != "0") ||
			(route["continuous_drop_off"] != "" && route["continuous_drop_off"] != "0") {
			hasContinuousPickupDropoff = true
			break
		}
	}

	// Process all stop times in a single pass
	for i, stopTime := range gtfs.Files["stop_times"] {
		stopTime, stopTimeMessages := parseStopTime(stopTime, gtfs.IdMap["trips"], gtfs.IdMap["stops"], gtfs.IdMap["location_groups"], hasContinuousPickupDropoff)
		stopTimes = append(stopTimes, stopTime)

		// Check for duplicate trip_id + stop_sequence combinations
		if stopTime.TripId != "" && stopTime.StopSequence > 0 {
			key := stopTime.TripId + "_" + strconv.Itoa(stopTime.StopSequence)
			if uniqueStopSequences[key] {
				messages = append(messages, types.Message{
					Field:        "stop_sequence",
					FileName:     "stop_times.txt",
					Message:      "Duplicate stop_sequence found for trip_id. Each stop_sequence must be unique within a trip.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			uniqueStopSequences[key] = true
		}

		// Update row number and other fields for each message
		for _, msg := range stopTimeMessages {
			msg.Row = i + 1
			msg.FileName = "stop_times.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}

	return stopTimes, messages
}

func parseStopTime(m map[string]string, tripIds map[string]int, stopIds map[string]int, locationGroupIds map[string]int, hasContinuousPickupDropoff bool) (stopTime types.StopTime, messages []types.Message) {
	// Pre-allocate messages slice with a reasonable capacity
	messages = make([]types.Message, 0, 10)

	// Use a single slice for parsing errors to avoid multiple allocations
	parsingErrors := make([]string, 0, 5)

	// Convert Optional Primitive Values
	var stopSequence int
	var pickupType, dropOffType int
	var continuousPickup, continuousDropOff int
	var timepoint bool
	var shapeDistTraveled float32
	var stopId, locationGroupId, locationId, stopHeadsign, startPickupDropOffWindow, endPickupDropOffWindow, pickupBookingRuleId, dropOffBookingRuleId, arrivalTime, departureTime string

	// Batch parse operations to reduce function call overhead
	lib.ParseStringToPrimitive(m["stop_sequence"], &stopSequence, &parsingErrors)
	lib.ParseStringToPrimitive(m["pickup_type"], &pickupType, &parsingErrors)
	lib.ParseStringToPrimitive(m["drop_off_type"], &dropOffType, &parsingErrors)
	lib.ParseStringToPrimitive(m["continuous_pickup"], &continuousPickup, &parsingErrors)
	lib.ParseStringToPrimitive(m["continuous_drop_off"], &continuousDropOff, &parsingErrors)
	lib.ParseStringToPrimitive(m["timepoint"], &timepoint, &parsingErrors)
	lib.ParseStringToPrimitive(m["shape_dist_traveled"], &shapeDistTraveled, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_id"], &stopId, &parsingErrors)
	lib.ParseStringToPrimitive(m["location_group_id"], &locationGroupId, &parsingErrors)
	lib.ParseStringToPrimitive(m["location_id"], &locationId, &parsingErrors)
	lib.ParseStringToPrimitive(m["stop_headsign"], &stopHeadsign, &parsingErrors)
	lib.ParseStringToPrimitive(m["start_pickup_drop_off_window"], &startPickupDropOffWindow, &parsingErrors)
	lib.ParseStringToPrimitive(m["end_pickup_drop_off_window"], &endPickupDropOffWindow, &parsingErrors)
	lib.ParseStringToPrimitive(m["pickup_booking_rule_id"], &pickupBookingRuleId, &parsingErrors)
	lib.ParseStringToPrimitive(m["drop_off_booking_rule_id"], &dropOffBookingRuleId, &parsingErrors)
	lib.ParseStringToPrimitive(m["arrival_time"], &arrivalTime, &parsingErrors)
	lib.ParseStringToPrimitive(m["departure_time"], &departureTime, &parsingErrors)

	// Assign values to stopTime struct
	stopTime.StopSequence = stopSequence
	stopTime.PickupType = lib.IfThenElse(m["pickup_type"] != "", &pickupType, nil)
	stopTime.DropOffType = lib.IfThenElse(m["drop_off_type"] != "", &dropOffType, nil)
	stopTime.ContinuousPickup = lib.IfThenElse(m["continuous_pickup"] != "", &continuousPickup, nil)
	stopTime.ContinuousDropOff = lib.IfThenElse(m["continuous_drop_off"] != "", &continuousDropOff, nil)
	stopTime.Timepoint = lib.IfThenElse(m["timepoint"] != "", &timepoint, nil)
	stopTime.ShapeDistTraveled = lib.IfThenElse(m["shape_dist_traveled"] != "", &shapeDistTraveled, nil)
	stopTime.StopId = stopId
	stopTime.LocationGroupId = lib.IfThenElse(m["location_group_id"] != "", &locationGroupId, nil)
	stopTime.LocationId = lib.IfThenElse(m["location_id"] != "", &locationId, nil)
	stopTime.StopHeadsign = lib.IfThenElse(m["stop_headsign"] != "", &stopHeadsign, nil)
	stopTime.StartPickupDropOffWindow = lib.IfThenElse(m["start_pickup_drop_off_window"] != "", &startPickupDropOffWindow, nil)
	stopTime.EndPickupDropOffWindow = lib.IfThenElse(m["end_pickup_drop_off_window"] != "", &endPickupDropOffWindow, nil)
	stopTime.PickupBookingRuleId = lib.IfThenElse(m["pickup_booking_rule_id"] != "", &pickupBookingRuleId, nil)
	stopTime.DropOffBookingRuleId = lib.IfThenElse(m["drop_off_booking_rule_id"] != "", &dropOffBookingRuleId, nil)
	stopTime.ArrivalTime = lib.IfThenElse(m["arrival_time"] != "", &arrivalTime, nil)
	stopTime.DepartureTime = lib.IfThenElse(m["departure_time"] != "", &departureTime, nil)

	// Convert Required Values
	lib.ParseStringToPrimitive(m["trip_id"], &stopTime.TripId, &parsingErrors)

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
	if stopTime.TripId == "" {
		messages = append(messages, types.Message{
			Field:   "trip_id",
			Message: "Trip ID is required.",
		})
	} else {
		_, ok := tripIds[stopTime.TripId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "trip_id",
				Message: "Trip ID must reference a valid trip_id from trips.txt.",
			})
		}
	}

	// Validate Required stop_sequence
	if stopTime.StopSequence < 0 {
		messages = append(messages, types.Message{
			Field:   "stop_sequence",
			Message: "Stop sequence is required and must be a non-negative integer.",
		})
	}

	// Validate stop_id, location_group_id, and location_id are mutually exclusive
	hasStopId := stopTime.StopId != ""
	hasLocationGroupId := stopTime.LocationGroupId != nil && *stopTime.LocationGroupId != ""
	hasLocationId := stopTime.LocationId != nil && *stopTime.LocationId != ""

	if hasStopId && hasLocationGroupId {
		messages = append(messages, types.Message{
			Field:   "stop_id",
			Message: "Stop ID and Location Group ID cannot both be defined.",
		})
	}

	if hasStopId && hasLocationId {
		messages = append(messages, types.Message{
			Field:   "stop_id",
			Message: "Stop ID and Location ID cannot both be defined.",
		})
	}

	if hasLocationGroupId && hasLocationId {
		messages = append(messages, types.Message{
			Field:   "location_group_id",
			Message: "Location Group ID and Location ID cannot both be defined.",
		})
	}

	// Validate at least one of stop_id, location_group_id, or location_id is defined
	if !hasStopId && !hasLocationGroupId && !hasLocationId {
		messages = append(messages, types.Message{
			Field:   "stop_id",
			Message: "At least one of Stop ID, Location Group ID, or Location ID must be defined.",
		})
	}

	// Validate stop_id if provided
	if hasStopId {
		_, ok := stopIds[stopTime.StopId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "stop_id",
				Message: "Stop ID must reference a valid stop_id from stops.txt.",
			})
		}
	}

	// Validate location_group_id if provided
	if hasLocationGroupId {
		_, ok := locationGroupIds[*stopTime.LocationGroupId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "location_group_id",
				Message: "Location Group ID must reference a valid location_group_id from location_groups.txt.",
			})
		}
	}

	// Validate arrival_time and departure_time are required for timepoint=1
	if stopTime.Timepoint != nil && *stopTime.Timepoint {
		if stopTime.ArrivalTime == nil || *stopTime.ArrivalTime == "" {
			messages = append(messages, types.Message{
				Field:   "arrival_time",
				Message: "Arrival time is required when timepoint=1.",
			})
		}
		if stopTime.DepartureTime == nil || *stopTime.DepartureTime == "" {
			messages = append(messages, types.Message{
				Field:   "departure_time",
				Message: "Departure time is required when timepoint=1.",
			})
		}
	}

	// Validate arrival_time and departure_time are forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if (stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") ||
		(stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "") {
		if stopTime.ArrivalTime != nil && *stopTime.ArrivalTime != "" {
			messages = append(messages, types.Message{
				Field:   "arrival_time",
				Message: "Arrival time is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.",
			})
		}
		if stopTime.DepartureTime != nil && *stopTime.DepartureTime != "" {
			messages = append(messages, types.Message{
				Field:   "departure_time",
				Message: "Departure time is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.",
			})
		}
	}

	// Validate start_pickup_drop_off_window and end_pickup_drop_off_window are required together
	if (stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") !=
		(stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "") {
		messages = append(messages, types.Message{
			Field:   "start_pickup_drop_off_window",
			Message: "Start pickup drop off window and end pickup drop off window must be defined together.",
		})
	}

	// Validate start_pickup_drop_off_window and end_pickup_drop_off_window are required when location_group_id or location_id is defined
	if (hasLocationGroupId || hasLocationId) &&
		((stopTime.StartPickupDropOffWindow == nil || *stopTime.StartPickupDropOffWindow == "") ||
			(stopTime.EndPickupDropOffWindow == nil || *stopTime.EndPickupDropOffWindow == "")) {
		messages = append(messages, types.Message{
			Field:   "start_pickup_drop_off_window",
			Message: "Start pickup drop off window and end pickup drop off window are required when location_group_id or location_id is defined.",
		})
	}

	// Validate pickup_type and drop_off_type enum values
	if stopTime.PickupType != nil {
		validPickupType := map[int]bool{0: true, 1: true, 2: true, 3: true}
		if !validPickupType[*stopTime.PickupType] {
			messages = append(messages, types.Message{
				Field:   "pickup_type",
				Message: "Invalid pickup_type value. Valid values are 0, 1, 2, 3.",
			})
		}

		// Validate pickup_type=0 or pickup_type=3 is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
		if (*stopTime.PickupType == 0 || *stopTime.PickupType == 3) &&
			((stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") ||
				(stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "")) {
			messages = append(messages, types.Message{
				Field:   "pickup_type",
				Message: "pickup_type=0 or pickup_type=3 is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.",
			})
		}
	}

	if stopTime.DropOffType != nil {
		validDropOffType := map[int]bool{0: true, 1: true, 2: true, 3: true}
		if !validDropOffType[*stopTime.DropOffType] {
			messages = append(messages, types.Message{
				Field:   "drop_off_type",
				Message: "Invalid drop_off_type value. Valid values are 0, 1, 2, 3.",
			})
		}

		// Validate drop_off_type=0 is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
		if *stopTime.DropOffType == 0 &&
			((stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") ||
				(stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "")) {
			messages = append(messages, types.Message{
				Field:   "drop_off_type",
				Message: "drop_off_type=0 is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.",
			})
		}
	}

	// Validate continuous_pickup and continuous_drop_off enum values
	if stopTime.ContinuousPickup != nil {
		validContinuousPickup := map[int]bool{0: true, 1: true, 2: true, 3: true}
		if !validContinuousPickup[*stopTime.ContinuousPickup] {
			messages = append(messages, types.Message{
				Field:   "continuous_pickup",
				Message: "Invalid continuous_pickup value. Valid values are 0, 1, 2, 3.",
			})
		}

		// Validate continuous_pickup is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
		if (stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") ||
			(stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "") {
			messages = append(messages, types.Message{
				Field:   "continuous_pickup",
				Message: "Continuous pickup is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.",
			})
		}
	}

	if stopTime.ContinuousDropOff != nil {
		validContinuousDropOff := map[int]bool{0: true, 1: true, 2: true, 3: true}
		if !validContinuousDropOff[*stopTime.ContinuousDropOff] {
			messages = append(messages, types.Message{
				Field:   "continuous_drop_off",
				Message: "Invalid continuous_drop_off value. Valid values are 0, 1, 2, 3.",
			})
		}

		// Validate continuous_drop_off is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
		if (stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") ||
			(stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "") {
			messages = append(messages, types.Message{
				Field:   "continuous_drop_off",
				Message: "Continuous drop off is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.",
			})
		}
	}

	// Validate timepoint enum values
	if stopTime.Timepoint != nil {
		// timepoint is a boolean in the struct, so we don't need to validate enum values
		// The value will be true for 1 and false for 0
	}

	// Validate pickup_booking_rule_id is recommended when pickup_type=2
	if stopTime.PickupType != nil && *stopTime.PickupType == 2 &&
		(stopTime.PickupBookingRuleId == nil || *stopTime.PickupBookingRuleId == "") {
		messages = append(messages, types.Message{
			Field:   "pickup_booking_rule_id",
			Message: "Pickup booking rule ID is recommended when pickup_type=2.",
		})
	}

	// Validate drop_off_booking_rule_id is recommended when drop_off_type=2
	if stopTime.DropOffType != nil && *stopTime.DropOffType == 2 &&
		(stopTime.DropOffBookingRuleId == nil || *stopTime.DropOffBookingRuleId == "") {
		messages = append(messages, types.Message{
			Field:   "drop_off_booking_rule_id",
			Message: "Drop off booking rule ID is recommended when drop_off_type=2.",
		})
	}

	return stopTime, messages
}
