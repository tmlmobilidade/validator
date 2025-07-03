package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_tariffs_information
  - Presence: Optional
  - Type: Enum

# Description

Describes if the stop has tariffs information.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasTariffsInformationValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasTariffsInformation.Severity != "" {
		s = rules.HasTariffsInformation.Severity
	}

	if s == types.SEVERITY_IGNORE {
		return
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

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasTariffsInformation) {
		addMessage("has_tariffs_information must be 0, 1, 2, or 3", types.SEVERITY_ERROR)
		return
	}

	// Validate value based on rules
	if rules != nil && rules.HasTariffsInformation.Options != nil {
		if slices.Contains(*rules.HasTariffsInformation.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasTariffsInformation.Options, strconv.Itoa(*stop.HasTariffsInformation)) {
			addMessage(fmt.Sprintf("has_tariffs_information is not allowed: %d", *stop.HasTariffsInformation), types.SEVERITY_ERROR)
			return
		}
	}
}
