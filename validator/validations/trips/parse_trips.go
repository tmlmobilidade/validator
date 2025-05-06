package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

// Parses a trip row from the trips.txt file to a Trip struct
func ParseTrips(rawTrips map[string]string, row int, gtfs *types.Gtfs) (trips types.Trip) {
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
		}
	}

	trips.TripId = tripId
	trips.RouteId = routeId
	trips.ServiceId = serviceId

	trips.TripHeadsign = lib.IfThenElse(tripHeadsign != "", &tripHeadsign, nil)
	trips.TripShortName = lib.IfThenElse(tripShortName != "", &tripShortName, nil)
	trips.BlockId = lib.IfThenElse(blockId != "", &blockId, nil)
	trips.ShapeId = lib.IfThenElse(shapeId != "", &shapeId, nil)
	trips.BikesAllowed = lib.IfThenElse(bikesAllowed != "", &bikesAllowedInt, nil)
	trips.WheelchairAccessible = lib.IfThenElse(wheelchairAccessible != "", &wheelchairAccessibleInt, nil)
	trips.DirectionId = lib.IfThenElse(directionId != "", &directionIdInt, nil)
	
	return trips
}

