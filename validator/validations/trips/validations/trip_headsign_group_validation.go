package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	"sort"
	"strings"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: trip_headsign
  - Presence: Optional (Required for "Transportes Metropolitanos de Lisboa" when pattern_id is set)
  - Type: Text

# Description

Validates that trip_headsign is unique within each pattern_id.
All trips with the same pattern_id must have the same trip_headsign (including consistent presence: nil/empty vs non-empty).
*/

func TripHeadsignGroupValidation(tripsGroupedByPattern types.TripGroupedByPattern, gtfs *types.Gtfs, rules *types.TripsRules) {

	for patternId, group := range tripsGroupedByPattern {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		headsigns := make(map[string]bool)
		for _, trip := range group.Trips {
			key := ""
			if trip.TripHeadsign != nil {
				key = *trip.TripHeadsign
			}
			headsigns[key] = true
		}

		if len(headsigns) <= 1 {
			continue
		}

		keys := make([]string, 0, len(headsigns))
		for h := range headsigns {
			keys = append(keys, h)
		}
		sort.Strings(keys)
		parts := make([]string, len(keys))
		for i, h := range keys {
			if h == "" {
				parts[i] = "[empty]"
			} else {
				parts[i] = h
			}
		}
		row := group.Trips[0].Row
		ctx := lib.NewValidationContext("trip_headsign", "trips.txt", "trip_headsign_group_validation", row, services.AppMessageService)
		if rules != nil && rules.TripHeadsignGroup.Severity != "" {
			ctx.WithSeverity(rules.TripHeadsignGroup.Severity)
		}
		if ctx.ShouldSkip() {
			return
		}
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("trip_headsign_group_validation.different_headsigns_in_pattern", patternId, strings.Join(parts, ", ")))
	}
}
