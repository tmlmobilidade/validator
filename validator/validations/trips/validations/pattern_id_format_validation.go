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
  - Field: pattern_id
  - Presence: optional (Required for "Transportes Metropolitanos de Lisboa")
  - Type: Foreigh Key referencing patterns.pattern_id

# Description

Validates if the pattern_id is in the correct format. Must be in the format "X at XXXX_X_[1/2/3]".
*/
func PatternIdFormatValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_format_validation", row, services.AppMessageService)
	if rules != nil && rules.PatternIdFormat.Severity != "" {
		ctx.WithSeverity(rules.PatternIdFormat.Severity)
	}

	if trip.PatternId == nil {
		return
	}

	patternIdParts := strings.Split(*trip.PatternId, "_")
	if len(patternIdParts) != 3 {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format_validation.invalid"))
		return
	}

	Parts1Valid := len(patternIdParts[0]) <= 4 && len(patternIdParts[0]) >= 1
	Parts2Valid := len(patternIdParts[1]) == 1

	Parts3Valid := false
	if len(patternIdParts[2]) == 1 {
		if n, err := strconv.Atoi(patternIdParts[2]); err == nil && (n == 1 || n == 2 || n == 3) {
			Parts3Valid = true
		}
	}

	if !Parts1Valid || !Parts2Valid || !Parts3Valid {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format_validation.invalid"))
		return
	}
}
