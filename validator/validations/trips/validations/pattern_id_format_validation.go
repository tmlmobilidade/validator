package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	"strings"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: pattern_id
  - Presence: optional (Required for "Transportes Metropolitanos de Lisboa")
  - Type: Foreigh Key referencing patterns.pattern_id

# Description

Validates if the pattern_id is in the correct format. Must be in the format "XXXX_[0|1]_X".
*/
func PatternIdFormatValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("pattern_id_format", "trips.txt", "pattern_id_format", row, services.AppMessageService)
	if rules != nil && rules.PatternIdFormat.Severity != "" {
		ctx.WithSeverity(rules.PatternIdFormat.Severity)
	}

	if trip.PatternId == nil {
		return
	}

	patternIdParts := strings.Split(*trip.PatternId, "_")
	if len(patternIdParts) != 3 {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format.invalid"))
		return
	}

	if len(patternIdParts[0]) != 4 || len(patternIdParts[1]) != 1 || len(patternIdParts[2]) != 1 {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format.invalid"))
		return
	}
}
