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

const defaultPatternIDFormat = "XXXX_X_X"

var patternIDFormatRegexMap = map[string]*regexp.Regexp{
	"XXXX_X_X":   regexp.MustCompile(`^[^_]{1,4}_[^_]_[^_]$`),
	"XXXX_X_XXX": regexp.MustCompile(`^[^_]{1,4}_[^_]_[^_]{3}$`),
}

func PatternIdFormatValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_format_validation", row, services.AppMessageService)
	if rules != nil && rules.PatternIdFormat.Severity != "" {
		ctx.WithSeverity(rules.PatternIdFormat.Severity)
	}

	if trip.PatternId == nil {
		return
	}

	expectedFormat := defaultPatternIDFormat
	patternIDRegex := patternIDFormatRegexMap[defaultPatternIDFormat]

	if rules != nil && rules.PatternIdFormat.Options != nil && len(*rules.PatternIdFormat.Options) > 0 {
		configuredFormat := (*rules.PatternIdFormat.Options)[0]
		if configuredRegex, ok := patternIDFormatRegexMap[configuredFormat]; ok {
			expectedFormat = configuredFormat
			patternIDRegex = configuredRegex
		}
	}

	if !patternIDRegex.MatchString(*trip.PatternId) {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format_validation.invalid", expectedFormat, *trip.PatternId))
		return
	}
}
