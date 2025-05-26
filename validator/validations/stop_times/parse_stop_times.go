package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseStopTimes(rawStopTimes map[string]string, row int) types.StopTime {
	var (
		stopTime types.StopTime = types.StopTime{}
		tripId,arrivalTime,departureTime,stopId,locationGroupStopId,locationId,stopHeadsign,startPickupDropOffWindow,endPickupDropOffWindow,pickupBookingRuleId,dropOffBookingRuleId string
		stopSequence, timepoint, pickupType, dropOffType, continuousPickup, continuousDropOff int
		shapeDistTraveled float64
		messages []types.Message
	)

	stringFields := map[string]*string{
		"stop_id": &stopId,
		"trip_id": &tripId,
		"arrival_time": &arrivalTime,
		"departure_time": &departureTime,
		"stop_headsign": &stopHeadsign,
		"start_pickup_drop_off_window": &startPickupDropOffWindow,
		"end_pickup_drop_off_window": &endPickupDropOffWindow,
		"pickup_booking_rule_id": &pickupBookingRuleId,
		"drop_off_booking_rule_id": &dropOffBookingRuleId,
	}

	intFields := map[string]*int{
		"stop_sequence": &stopSequence,
		"pickup_type": &pickupType,
		"drop_off_type": &dropOffType,
		"continuous_pickup": &continuousPickup,
		"continuous_drop_off": &continuousDropOff,
		"timepoint": &timepoint,
	}

	floatFields := map[string]*float64{
		"shape_dist_traveled": &shapeDistTraveled,
	}

	// Helper to collect error messages
	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "stop_times.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "stop_times_parse",
		})
	}

	// Parse string fields
	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawStopTimes[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	// Parse int fields
	for field, target := range intFields {
		if errMsg := lib.ParseStringToPrimitive(rawStopTimes[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	// Parse float fields
	for field, target := range floatFields {
		if errMsg := lib.ParseStringToPrimitive(rawStopTimes[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return types.StopTime{}
	}

	stopTime.TripId = lib.IfThenElse(rawStopTimes["trip_id"] != "", &tripId, nil)
	stopTime.ArrivalTime = lib.IfThenElse(rawStopTimes["arrival_time"] != "", &arrivalTime, nil)
	stopTime.DepartureTime = lib.IfThenElse(rawStopTimes["departure_time"] != "", &departureTime, nil)
	stopTime.StopId = lib.IfThenElse(rawStopTimes["stop_id"] != "", &stopId, nil)
	stopTime.LocationGroupStopId = lib.IfThenElse(rawStopTimes["location_group_stop_id"] != "", &locationGroupStopId, nil)
	stopTime.LocationId = lib.IfThenElse(rawStopTimes["location_id"] != "", &locationId, nil)
	stopTime.StopSequence = lib.IfThenElse(rawStopTimes["stop_sequence"] != "", &stopSequence, nil)
	stopTime.StopHeadsign = lib.IfThenElse(rawStopTimes["stop_headsign"] != "", &stopHeadsign, nil)
	stopTime.StartPickupDropOffWindow = lib.IfThenElse(rawStopTimes["start_pickup_drop_off_window"] != "", &startPickupDropOffWindow, nil)
	stopTime.EndPickupDropOffWindow = lib.IfThenElse(rawStopTimes["end_pickup_drop_off_window"] != "", &endPickupDropOffWindow, nil)
	stopTime.PickupType = lib.IfThenElse(rawStopTimes["pickup_type"] != "", &pickupType, nil)
	stopTime.DropOffType = lib.IfThenElse(rawStopTimes["drop_off_type"] != "", &dropOffType, nil)
	stopTime.ContinuousPickup = lib.IfThenElse(rawStopTimes["continuous_pickup"] != "", &continuousPickup, nil)
	stopTime.ContinuousDropOff = lib.IfThenElse(rawStopTimes["continuous_drop_off"] != "", &continuousDropOff, nil)
	stopTime.ShapeDistTraveled = lib.IfThenElse(rawStopTimes["shape_dist_traveled"] != "", &shapeDistTraveled, nil)
	stopTime.Timepoint = lib.IfThenElse(rawStopTimes["timepoint"] != "", &timepoint, nil)
	stopTime.PickupBookingRuleId = lib.IfThenElse(rawStopTimes["pickup_booking_rule_id"] != "", &pickupBookingRuleId, nil)
	stopTime.DropOffBookingRuleId = lib.IfThenElse(rawStopTimes["drop_off_booking_rule_id"] != "", &dropOffBookingRuleId, nil)
	
	return stopTime
}