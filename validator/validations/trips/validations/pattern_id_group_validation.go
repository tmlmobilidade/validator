package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: trips.txt
  - Field: pattern_id
  - Presence: Optional (Required for "Transportes Metropolitanos de Lisboa")
  - Type: ID

# Description

Validates if trips with the same pattern_id have the same route_id, trip_headsign, direction_id, shape_id and the same stop sequence.
*/
func PatternIdGroupValidation(tripsGroupedByPattern types.TripGroupedByPattern, gtfs *types.Gtfs, rules *types.TripsRules) {
	// Group trips by pattern_id and validate the group
	for patternId, group := range tripsGroupedByPattern {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		for _, trip := range group.Trips {
			ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_validation", trip.Row, services.AppMessageService)
			if rules != nil && rules.PatternIdGroup.Severity != "" {
				ctx.WithSeverity(rules.PatternIdGroup.Severity)
			}
			if ctx.ShouldSkip() {
				return
			}
			if trip.ShapeId == nil {
				ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("pattern_id_validation.shape_id_not_found"))
				continue
			}

			if trip.RouteId == nil {
				ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("pattern_id_validation.route_id_not_found"))
				continue
			}

			if trip.DirectionId == nil {
				ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("pattern_id_validation.direction_id_not_found"))
				continue
			}

			if trip.TripHeadsign == nil {
				ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("pattern_id_validation.trip_headsign_not_found"))
				continue
			}
		}

		if len(group.Hash) > 1 {
			row := group.Trips[0].Row
			ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_validation", row, services.AppMessageService)
			if rules != nil && rules.PatternIdGroup.Severity != "" {
				ctx.WithSeverity(rules.PatternIdGroup.Severity)
			}
			if ctx.ShouldSkip() {
				return
			}
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("pattern_id_validation.multiple_stop_sequence_variations", patternId))
		}
	}
}
