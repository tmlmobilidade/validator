package agency

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyIdValidation_Required(t *testing.T) {
	rules := &types.GtfsRules{Agency: types.AgencyRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}}
	agency := &types.Agency{AgencyId: nil}
	gtfs := types.Gtfs{Agency: []types.AgencyRaw{{}, {}}}
	validations.AgencyIdValidation(agency, 1, gtfs, &rules.Agency)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Agency ID is required when there is more than one agency",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Recommended(t *testing.T) {
	rules := &types.GtfsRules{Agency: types.AgencyRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_WARNING}}}
	agency := &types.Agency{AgencyId: nil}
	gtfs := types.Gtfs{Agency: []types.AgencyRaw{{}}}
	validations.AgencyIdValidation(agency, 2, gtfs, &rules.Agency)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Agency ID is recomended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAllIdHelpers(t *testing.T) {
	for _, tc := range test_helpers.GetIdTestCases() {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			validations.AgencyIdValidation(tc.Agency, tc.Row, *tc.Gtfs, nil)

			assertion := lib.AssertionMessage{
				Expected: tc.ExpectedErrors,
				Actual:   services.AppMessageService.GetSummary().TotalErrors,
				Message:  tc.Name,
			}
			if assert := lib.Assert(assertion); assert != "" {
				t.Error(assert)
			}
		})
	}
}
