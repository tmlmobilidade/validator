package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_tariffs_information
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has tariffs information.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasTariffsInformationValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasTariffsInformation.Severity != "" {
		s = rules.HasTariffsInformation.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_tariffs_information",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_tariffs_information_validation",
		})
	}

	if stop.HasTariffsInformation == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_tariffs_information is required", "has_tariffs_information is recommended")
		addMessage(warn, s)
		return
	}
}
