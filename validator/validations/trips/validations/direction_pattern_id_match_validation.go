package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	"strconv"
	"strings"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: direction_id
  - Presence: Required
  - Type: ID

# Description

Ensure the direction_id is consistent with the pattern_id (e.g., pattern_id "1001_0_1" should have direction_id = 0).
*/
func DirectionPatternIdMatchValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("direction_pattern_id_match", "trips.txt", "direction_pattern_id_match", row, services.AppMessageService)

	// Use DirectionPatternIdMatch severity if available, otherwise fallback to DirectionId severity
	if rules != nil {
		if rules.DirectionPatternIdMatch.Severity != "" {
			ctx.WithSeverity(rules.DirectionPatternIdMatch.Severity)
		} else if rules.DirectionId.Severity != "" {
			ctx.WithSeverity(rules.DirectionId.Severity)
		}
	}

	if ctx.ShouldSkip() {
		return
	}

	// Handle required fields
	if trip.PatternId == nil || trip.DirectionId == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("direction_pattern_id_match.required"))
		return
	}

	// Validate pattern_id format: XXXX_X_X (6 characters total: 4 chars + _ + 1 char + _ + 1 char)
	patternIdParts := strings.Split(*trip.PatternId, "_")
	if len(*trip.PatternId) != 8 || len(patternIdParts) != 3 {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("direction_pattern_id_match.invalid_pattern_id"))
		return
	}

	// Validate each part: first part should be 4 characters, second and third should be 1 character each
	if len(patternIdParts[0]) != 4 || len(patternIdParts[1]) != 1 || len(patternIdParts[2]) != 1 {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("direction_pattern_id_match.invalid_pattern_id"))
		return
	}

	// Parse the directionId part (second part) as integer
	directionId, err := strconv.Atoi(patternIdParts[1])
	if err != nil || directionId < 0 || directionId > 1 {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("direction_pattern_id_match.invalid_direction_id"))
		return
	}

	// Ensure trip.DirectionId matches parsed directionId from patternId
	if *trip.DirectionId != directionId {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("direction_pattern_id_match.not_matching"))
		return
	}
}
