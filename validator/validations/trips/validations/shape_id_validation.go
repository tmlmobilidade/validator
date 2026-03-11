package trips

import (
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
func ShapeIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules, tripStopTimesCache map[string][]types.StopTimeRaw) {
	ctx := lib.NewValidationContext("shape_id", "trips.txt", "shape_id_validation", row, services.AppMessageService)
	if rules != nil && rules.ShapeId.Severity != "" {
		ctx.WithSeverity(rules.ShapeId.Severity)
	}

	// Never let a malformed row crash the full validator process.
	defer func() {
		if panicValue := recover(); panicValue != nil {
			ctx.AddError(ctx.GetTranslatedMessage("shape_id_validation.internal_error", panicValue))
		}
	}()

	hasContinuousPickupDropoff := false

	if trip == nil || trip.RouteId == nil {
		return
	}

	// Check if the route has continuous pickup/dropoff behavior
	routeRows, err := gtfs.GetRowsById("routes", *trip.RouteId)
	if err != nil {
		ctx.AddError(ctx.GetTranslatedMessage("shape_id_validation.route_lookup_failed", *trip.RouteId, err.Error()))
		return
	}
	if len(routeRows) == 0 {
		ctx.AddError(ctx.GetTranslatedMessage("shape_id_validation.route_not_found_for_trip", *trip.RouteId))
		return
	}

	for _, routeRow := range routeRows {
		routeRaw, getRouteErr := gtfs.GetRoute(routeRow)
		if getRouteErr != nil {
			continue
		}
		if routeRaw.ContinuousPickup != "" || routeRaw.ContinuousDropOff != "" {
			hasContinuousPickupDropoff = true
		}
		break
	}

	// Check if the stop_times have continuous pickup/dropoff behavior
	// Use cached stop_times data instead of querying database
	if trip.TripId != nil {
		stopTimesRaw, exists := tripStopTimesCache[*trip.TripId]
		if exists && !hasContinuousPickupDropoff {
			for _, stopTimeRaw := range stopTimesRaw {
				if stopTimeRaw.ContinuousPickup != "" || stopTimeRaw.ContinuousDropOff != "" {
					hasContinuousPickupDropoff = true
					break // Exit early once we find a continuous pickup
				}
			}
		} else if !exists {
			// Fallback to database query if not in cache (shouldn't happen)
			stopTimeRows, err := gtfs.GetRowsById("stop_times", *trip.TripId)
			if err == nil && len(stopTimeRows) > 0 && !hasContinuousPickupDropoff {
				for _, rowIndex := range stopTimeRows {
					stopTimeRaw, err := gtfs.GetStopTime(rowIndex)
					if err != nil {
						continue
					}
					if stopTimeRaw.ContinuousPickup != "" || stopTimeRaw.ContinuousDropOff != "" {
						hasContinuousPickupDropoff = true
						break // Exit early once we find a continuous pickup
					}
				}
			}
		}
	}

	if hasContinuousPickupDropoff && trip.ShapeId == nil {
		ctx.AddError(ctx.GetTranslatedMessage("shape_id_validation.continuous_pickup_dropoff"))
		return
	}

	if trip.ShapeId == nil {
		if ctx.ShouldSkip() {
			return
		}

		message := ctx.GetRequiredMessage("shape_id_validation.required", "shape_id_validation.recommended")
		ctx.AddMessageWithSeverity(message)
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(gtfs, "shapes", *trip.ShapeId) {
		ctx.AddError(ctx.GetTranslatedMessage("shape_id_validation.not_found", *trip.ShapeId))
		return
	}
}
