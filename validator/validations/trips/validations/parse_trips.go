package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

// Parses a trip row from the trips.txt file to a Trip struct
func ParseTrips(rawTrips map[string]string, row int, gtfs *types.Gtfs) (trip types.Trip) {
	message := types.Message{
		Field:   "",
		FileName: "trips.txt",
		Rows: []int{row},
		Message: "",
		Severity: types.SEVERITY_ERROR,
		ValidationID: "trips_parse",
	}

	var tripId, routeId, serviceId, tripHeadsign, tripShortName, directionId, blockId, shapeId, wheelchairAccessible, bikesAllowed string
	var directionIdInt, wheelchairAccessibleInt, bikesAllowedInt int

	fieldMappings := map[string]*string{
		"trip_id":               &tripId,
		"route_id":              &routeId,
		"service_id":            &serviceId,
		"trip_headsign":         &tripHeadsign,
		"trip_short_name":       &tripShortName,
		"direction_id":          &directionId,
		"block_id":              &blockId,
		"shape_id":              &shapeId,
		"wheelchair_accessible": &wheelchairAccessible,
		"bikes_allowed":         &bikesAllowed,
	}

	// Loop through fields and parse each one
	for field, target := range fieldMappings {
		msg := lib.ParseStringToPrimitive(rawTrips[field], target)
		if msg != "" {
			message.Message = msg
			message.Field = field
			services.AppMessageService.AddMessage(message)
			return types.Trip{}
		}
	}

	trip.TripId = tripId
	trip.RouteId = routeId
	trip.ServiceId = serviceId

	trip.TripHeadsign = lib.IfThenElse(tripHeadsign != "", &tripHeadsign, nil)
	trip.TripShortName = lib.IfThenElse(tripShortName != "", &tripShortName, nil)
	trip.BlockId = lib.IfThenElse(blockId != "", &blockId, nil)
	trip.ShapeId = lib.IfThenElse(shapeId != "", &shapeId, nil)
	trip.BikesAllowed = lib.IfThenElse(bikesAllowed != "", &bikesAllowedInt, nil)
	trip.WheelchairAccessible = lib.IfThenElse(wheelchairAccessible != "", &wheelchairAccessibleInt, nil)
	trip.DirectionId = lib.IfThenElse(directionId != "", &directionIdInt, nil)
	
	return trip
}

