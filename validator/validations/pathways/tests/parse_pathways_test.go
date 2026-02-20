package pathways_tests

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/pathways/validations"
	"testing"
)

func TestAllParsePathwaysTestCases(t *testing.T) {
	t.Run("Valid_All_Fields", func(t *testing.T) {
		rawPathways := types.PathwaysRaw{
			FromStopId:           "1",
			ToStopId:             "2",
			PathwayId:            "3",
			SignpostedAs:         "4",
			ReversedSignpostedAs: "5",
			MaxSlope:             "6",
			MinWidth:             "7",
			PathwayMode:          "8",
			IsBidirectional:      "9",
			TraversalTime:        "10",
			Length:               "11",
			StairCount:           "12",
		}
		validations.ParsePathways(rawPathways, 1)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Valid_All_Fields", types.SEVERITY_ERROR)
	})
	t.Run("Invalid_All_Fields", func(t *testing.T) {
		rawPathways := types.PathwaysRaw{
			FromStopId:           "1",
			ToStopId:             "not_a_stop_id",
			PathwayId:            "not_a_pathway_id",
			SignpostedAs:         "not_a_signposted_as",
			ReversedSignpostedAs: "not_a_reversed_signposted_as",
			MaxSlope:             "not_a_max_slope",
			MinWidth:             "not_a_min_width",
			TraversalTime:        "not_a_traversal_time",
		}
		validations.ParsePathways(rawPathways, 1)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Invalid_All_Fields", types.SEVERITY_ERROR)
	})
}
