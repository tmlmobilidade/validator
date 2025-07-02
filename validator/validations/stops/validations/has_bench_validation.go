package stops

import (
	"main/lib"
	"main/services"
	"main/types"
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
}
