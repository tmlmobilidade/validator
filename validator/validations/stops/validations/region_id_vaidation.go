package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: region_id
  - Presence: Optional
  - Type: String

# Description

Region identifier for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func RegionIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RegionId.Severity != "" {
		s = rules.RegionId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "region_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "region_id_validation",
		})
	}

	if stop.RegionId == nil || *stop.RegionId == "" {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"region_id_validation.required",
				"region_id_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("region_id_validation.forbidden"), s)
		return
	}

	// Validate rules
	if rules != nil && rules.RegionId.Options != nil {
		if slices.Contains(*rules.RegionId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RegionId.Options, *stop.RegionId) {
			addMessage(i18n.AppTranslator.Get("region_id_validation.not_allowed", *stop.RegionId), s)
			return
		}
	}
}
