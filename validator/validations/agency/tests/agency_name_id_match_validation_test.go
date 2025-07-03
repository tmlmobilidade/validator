package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

var compare = []types.Compare{
	{
		Key:   "1",
		Value: "Name 1",
	},
	{
		Key:   "2",
		Value: "Name 2",
	},
	{
		Key:   "3",
		Value: "Name 3",
	},
	{
		Key:   "4",
		Value: "Name 1",
	},
}

func TestAgencyNameIdMatchValidation_Valid(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyName: lib.Ptr("Name 1"), AgencyId: lib.Ptr("1")}
	rules := &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{
		Severity: severity,
		Compare:  &compare,
	}}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "No errors should be found",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyNameIdMatchValidation_Invalid(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyName: lib.Ptr("Name 4"), AgencyId: lib.Ptr("2")}
	rules := &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{
		Severity: severity,
		Compare:  &compare,
	}}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Should have found one error for agency id and name mismatch",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyNameIdMatchValidation_Ignore(t *testing.T) {
	agency := &types.Agency{AgencyName: lib.Ptr("Name 4"), AgencyId: lib.Ptr("2")}
	rules := &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{
		Severity: types.SEVERITY_IGNORE,
		Compare:  &compare,
	}}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "No errors should be found",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyNameIdMatchValidation_Valid_RulesDuplicate(t *testing.T) {
	agency := &types.Agency{AgencyName: lib.Ptr("Name 3"), AgencyId: lib.Ptr("3")}
	rules := &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{
		Severity: types.SEVERITY_ERROR,
		Compare: &[]types.Compare{
			{
				Key:   "1",
				Value: "Name 1",
			},
			{
				Key:   "2",
				Value: "Name 2",
			},
			{
				Key:   "1",
				Value: "Name 1",
			},
		},
	}}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "No errors should be found",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}
