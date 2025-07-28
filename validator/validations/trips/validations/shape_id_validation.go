package trips

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: shape_id
  - Presence: Conditionally Required
  - Type: Foreign Key referencing shapes.shape_id

# Description

Identifies a geospatial shape describing the vehicle travel path for a trip.

Conditionally Required:
  - Required if the trip has a continuous pickup or drop-off behavior defined either in routes.txt or in stop_times.txt.
  - Optional otherwise.

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func ShapeIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ShapeId.Severity != "" {
		s = rules.ShapeId.Severity
	}

	addMessage := func(message string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shape_id",
			FileName:     "trips.txt",
			Rows:         []int{row},
			Severity:     severity,
			Message:      message,
			ValidationID: "shape_id_validation",
		})
	}

	hasContinuousPickupDropoff := false

	if trip.RouteId == nil {
		return
	}

	// Check if the route has continuous pickup/dropoff behavior
	routeRow := gtfs.IdMap["routes"][*trip.RouteId]
	if gtfs.IdMap["routes"] != nil && gtfs.Route[routeRow[0]].ContinuousPickup != "" {
		hasContinuousPickupDropoff = true
	}

	// Check if the stop_times have continuous pickup/dropoff behavior
	if gtfs.IdMap["stop_times"] != nil && len(gtfs.IdMap["stop_times"][*trip.TripId]) > 0 && !hasContinuousPickupDropoff {

		for _, rowIndex := range gtfs.IdMap["stop_times"][*trip.TripId] {
			if continuousPickup := gtfs.StopTime[rowIndex].ContinuousPickup; continuousPickup != "" {
				hasContinuousPickupDropoff = true
				break // Exit early once we find a continuous pickup
			}
		}
	}

	if hasContinuousPickupDropoff && trip.ShapeId == nil {
		addMessage(i18n.AppTranslator.Get("shape_id_validation.required_with_continuous"), types.SEVERITY_ERROR)
		return
	}

	if trip.ShapeId == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("shape_id_validation.required"), i18n.AppTranslator.Get("shape_id_validation.recommended"))
		addMessage(message, s)
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(gtfs, "shapes", *trip.ShapeId) {
		addMessage(i18n.AppTranslator.Get("shape_id_validation.not_found", map[string]interface{}{"shape_id": *trip.ShapeId}), types.SEVERITY_ERROR)
		return
	}
}
