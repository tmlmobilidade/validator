package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	"regexp"
	"strings"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: pattern_id
  - Presence: optional (Required for "Transportes Metropolitanos de Lisboa")
  - Type: Foreigh Key referencing patterns.pattern_id

# Description

Validates if the pattern_id is in the correct format. Must be in the defalut format "X at XXXX_X_X" or the formats specified in the rules.
*/

const defaultPatternIDFormat = "XXXX_X_X"

var patternIDFormatRegexMap = map[string]*regexp.Regexp{
	"XXXX_X_X":    regexp.MustCompile(`^[^_]{1,4}_[^_]_[^_]$`),
	"XXXX_X_ASC":  regexp.MustCompile(`^[^_]{1,4}_[^_]_ASC$`),
	"XXXX_X_DESC": regexp.MustCompile(`^[^_]{1,4}_[^_]_DESC$`),
	"XXXX_X_CIRC": regexp.MustCompile(`^[^_]{1,4}_[^_]_CIRC$`),
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
	allowedRegexes := []*regexp.Regexp{patternIDFormatRegexMap[defaultPatternIDFormat]}

	if rules != nil && rules.PatternIdFormat.Options != nil && len(*rules.PatternIdFormat.Options) > 0 {
		expectedFormats := make([]string, 0)
		allowedRegexes = make([]*regexp.Regexp, 0)

		for _, configuredFormat := range *rules.PatternIdFormat.Options {
			if configuredRegex, ok := patternIDFormatRegexMap[configuredFormat]; ok {
				expectedFormats = append(expectedFormats, configuredFormat)
				allowedRegexes = append(allowedRegexes, configuredRegex)
			}
		}

		// Fallback to default when options are present but none are recognized.
		if len(allowedRegexes) == 0 {
			allowedRegexes = []*regexp.Regexp{patternIDFormatRegexMap[defaultPatternIDFormat]}
		} else {
			expectedFormat = strings.Join(expectedFormats, " | ")
		}
	}

	isValidFormat := false
	for _, allowedRegex := range allowedRegexes {
		if allowedRegex.MatchString(*trip.PatternId) {
			isValidFormat = true
			break
		}
	}

	if !isValidFormat {
		ctx.AddError(ctx.GetTranslatedMessage("pattern_id_format_validation.invalid", expectedFormat, *trip.PatternId))
		return
	}
}
