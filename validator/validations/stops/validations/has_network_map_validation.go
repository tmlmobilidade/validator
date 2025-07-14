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
  - Field: has_network_map
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a network map.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasNetworkMapValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasNetworkMap.Severity != "" {
		s = rules.HasNetworkMap.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_network_map",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_network_map_validation",
		})
	}

	if stop.HasNetworkMap == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"has_network_map_validation.required",
				"has_network_map_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("has_network_map_validation.forbidden"), s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasNetworkMap) {
		addMessage(i18n.AppTranslator.Get("has_network_map_validation.invalid", *stop.HasNetworkMap), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasNetworkMap.Options != nil {
		if slices.Contains(*rules.HasNetworkMap.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasNetworkMap.Options, strconv.Itoa(*stop.HasNetworkMap)) {
			addMessage(i18n.AppTranslator.Get("has_network_map_validation.not_allowed", *stop.HasNetworkMap), s)
			return
		}
	}
}
