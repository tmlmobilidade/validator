package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseTrips(rawTrips map[string]string, row int, gtfs *types.Gtfs) types.Trip {
	var (
		trip                                           types.Trip = types.Trip{}
		tripId, routeId, serviceId                     string
		tripHeadsign, tripShortName, blockId, shapeId  string
		directionId, wheelchairAccessible, bikesAllowed int
		messages                                       []types.Message
	)

	stringFields := map[string]*string{
		"trip_id":         &tripId,
		"route_id":        &routeId,
		"service_id":      &serviceId,
		"trip_headsign":   &tripHeadsign,
		"trip_short_name": &tripShortName,
		"block_id":        &blockId,
		"shape_id":        &shapeId,
	}

	intFields := map[string]*int{
		"direction_id":          &directionId,
		"wheelchair_accessible": &wheelchairAccessible,
		"bikes_allowed":         &bikesAllowed,
	}

	// Helper to collect error messages
	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "trips.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "trips_parse",
		})
	}

	// Parse string fields
	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawTrips[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	// Parse int fields
	for field, target := range intFields {
		if errMsg := lib.ParseStringToPrimitive(rawTrips[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	// If there are any errors, return an empty trip
	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return trip
	}

	// Required fields
	trip.TripId = tripId
	trip.RouteId = routeId
	trip.ServiceId = serviceId

	// Optional fields
	trip.TripHeadsign = lib.IfThenElse(rawTrips["trip_headsign"] != "", &tripHeadsign, nil)
	trip.TripShortName = lib.IfThenElse(rawTrips["trip_short_name"] != "", &tripShortName, nil)
	trip.BlockId = lib.IfThenElse(rawTrips["block_id"] != "", &blockId, nil)
	trip.ShapeId = lib.IfThenElse(rawTrips["shape_id"] != "", &shapeId, nil)
	trip.DirectionId = lib.IfThenElse(rawTrips["direction_id"] != "", &directionId, nil)
	trip.WheelchairAccessible = lib.IfThenElse(rawTrips["wheelchair_accessible"] != "", &wheelchairAccessible, nil)
	trip.BikesAllowed = lib.IfThenElse(rawTrips["bikes_allowed"] != "", &bikesAllowed, nil)

	return trip
}