package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: wheelchair_boarding
  - Presence: Optional
  - Type: Enum

# Description

Indicates whether wheelchair boardings are possible from the location.

Valid options are:

For parentless stops:

  - 0 or empty - No accessibility information for the stop.
  - 1 - Some vehicles at this stop can be boarded by a rider in a wheelchair.
  - 2 - Wheelchair boarding is not possible at this stop.

For child stops:

  - 0 or empty - Stop will inherit its wheelchair_boarding behavior from the parent station, if specified in the parent.
  - 1 - There exists some accessible path from outside the station to the specific stop/platform.
  - 2 - There exists no accessible path from outside the station to the specific stop/platform.

For station entrances/exits:

  - 0 or empty - Station entrance will inherit its wheelchair_boarding behavior from the parent station, if specified for the parent.
  - 1 - Station entrance is wheelchair accessible.
  - 2 - No accessible path from station entrance to stops/platforms.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func WheelchairBoardingValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.WheelchairBoarding.Severity != "" {
		s = rules.WheelchairBoarding.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "wheelchair_boarding",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "wheelchair_boarding_validation",
		})
	}

	// Validate presence
	if stop.WheelchairBoarding == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"wheelchair_boarding_validation.required",
				"wheelchair_boarding_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("wheelchair_boarding_validation.forbidden"), s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2}
	if !slices.Contains(validValues, *stop.WheelchairBoarding) {
		addMessage(i18n.AppTranslator.Get("wheelchair_boarding_validation.invalid", *stop.WheelchairBoarding), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.WheelchairBoarding.Options != nil {
		if slices.Contains(*rules.WheelchairBoarding.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.WheelchairBoarding.Options, strconv.Itoa(*stop.WheelchairBoarding)) {
			addMessage(i18n.AppTranslator.Get("wheelchair_boarding_validation.not_allowed", *stop.WheelchairBoarding), s)
			return
		}
	}
}
