package agency

import (
	"main/lib"
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
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency ID is required when there is more than one agency",
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
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency ID is recomended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Unique(t *testing.T) {
	id := "unique"
	agency := &types.Agency{AgencyId: &id}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"agency": {"unique": {1}}}}
	validations.AgencyIdValidation(agency, 3, gtfs, nil)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency ID should be unique",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Duplicate(t *testing.T) {
	id := "duplicate"
	agency := &types.Agency{AgencyId: &id}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"agency": {"duplicate": {1, 2}}}}
	validations.AgencyIdValidation(agency, 4, gtfs, nil)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency ID should be duplicate",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 