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
  - Field: has_bench
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a bench.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasBenchValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasBench.Severity != "" {
		s = rules.HasBench.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_bench",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_bench_validation",
		})
	}

	if stop.HasBench == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_bench is required", "has_bench is recommended")
		addMessage(warn, s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasBench) {
		addMessage("has_bench must be 0, 1, 2, or 3", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasBench.Options != nil {
		if slices.Contains(*rules.HasBench.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasBench.Options, strconv.Itoa(*stop.HasBench)) {
			addMessage(fmt.Sprintf("has_bench is not allowed: %d", *stop.HasBench), s)
		}

		return
	}
}
