package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	"regexp"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: pattern_id
  - Presence: optional (Required for "Transportes Metropolitanos de Lisboa")
  - Type: Foreigh Key referencing patterns.pattern_id

# Description

Validates if the pattern_id is in the correct format. Must be in the format "X at XXXX_X_X".
*/

var patternIDRegex = regexp.MustCompile(`^[^_]{1,4}_[^_]_[^_]$`)

func PatternIdFormatValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_format_validation", "pattern_id_matches_feed_pattern_id_syntax", row, services.AppMessageService)
	if rules != nil && rules.PatternIdFormat.Severity != "" {
		ctx.WithSeverity(rules.PatternIdFormat.Severity)
	}

	if trip.PatternId == nil {
		return
	}

	if !patternIDRegex.MatchString(*trip.PatternId) {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format_validation.invalid", *trip.PatternId))
		return
	}
}
