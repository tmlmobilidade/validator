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

1) All trips with the same pattern_id must have the same trip_headsign
   (including consistent presence: nil/empty vs non-empty).
2) The same non-empty trip_headsign must not appear on more than one pattern_id
   (duplicates are usually placeholders such as "to be defined").
   Empty or missing headsigns are not compared across pattern_ids.
*/

func tripHeadsignKey(trip types.Trip) string {
	if trip.TripHeadsign != nil {
		return *trip.TripHeadsign
	}
	return ""
}

func TripHeadsignGroupValidation(tripsGroupedByPattern types.TripGroupedByPattern, gtfs *types.Gtfs, rules *types.TripsRules) {

	// 1) All trips with the same pattern_id must have the same trip_headsign
	for patternId, group := range tripsGroupedByPattern {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		headsigns := make(map[string]bool)
		for _, trip := range group.Trips {
			headsigns[tripHeadsignKey(trip)] = true
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

	// 2) The same non-empty trip_headsign must not appear on more than one pattern_id
	headsignToPatternIds := make(map[string]map[string]struct{})
	for patternId, group := range tripsGroupedByPattern {
		seen := make(map[string]struct{})
		for _, trip := range group.Trips {
			k := tripHeadsignKey(trip)
			if k == "" {
				continue
			}
			seen[k] = struct{}{}
		}
		for k := range seen {
			if headsignToPatternIds[k] == nil {
				headsignToPatternIds[k] = make(map[string]struct{})
			}
			headsignToPatternIds[k][patternId] = struct{}{}
		}
	}

	for headsign, patternSet := range headsignToPatternIds {
		if len(patternSet) <= 1 {
			continue
		}
		pids := make([]string, 0, len(patternSet))
		for p := range patternSet {
			pids = append(pids, p)
		}
		sort.Strings(pids)
		manyPatterns := len(pids) > 100
		for _, patternId := range pids {
			g := tripsGroupedByPattern[patternId]
			for _, trip := range g.Trips {
				if tripHeadsignKey(trip) != headsign {
					continue
				}
				row := trip.Row
				ctx := lib.NewValidationContext("trip_headsign", "trips.txt", "trip_headsign_group_validation", row, services.AppMessageService)
				if rules != nil && rules.TripHeadsignGroup.Severity != "" {
					ctx.WithSeverity(rules.TripHeadsignGroup.Severity)
				}
				if ctx.ShouldSkip() {
					return
				}
				if manyPatterns {
					ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("trip_headsign_group_validation.same_headsign_on_multiple_pattern_ids_many", headsign))
				} else {
					ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("trip_headsign_group_validation.same_headsign_on_multiple_pattern_ids", headsign, patternId))
				}
			}
		}
	}
}
