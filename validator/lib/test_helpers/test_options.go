package test_helpers

import (
	"main/types"
	"strings"
	"testing"
)

type OptionTestCase struct {
	Name           string
	Gtfs           *types.Gtfs
	Agency         *types.Agency
	Options        []string
	Row            int
	ExpectedErrors int
}

func GetOptionTestCases() []OptionTestCase {
	return []OptionTestCase{
		{
			Name:           "Valid",
			Gtfs:           CreateTestGtfsWithValidOptions("agency", []string{"valid"}, 1),
			Options:        []string{"valid"},
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Invalid",
			Gtfs:           CreateTestGtfsWithInvalidOptions("agency", []string{"invalid"}, 1),
			Options:        []string{"invalid"},
			Row:            1,
			ExpectedErrors: 1,
		},
	}
}

func CreateTestGtfsWithValidOptions(entityType string, options []string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			entityType: {
				strings.Join(options, ","): {row},
			},
		},
	}
}

func CreateTestGtfsWithInvalidOptions(entityType string, options []string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			entityType: {
				strings.Join(options, ","): {row},
			},
		},
	}
}

func TestAllOptionHelpers(t *testing.T) {
	t.Helper()

	tests := []struct {
		name           string
		entityType     string
		options        []string
		row            int
		expectedErrors int
		createFunc     func(string, []string, int) *types.Gtfs
	}{
		{
			name:           "Valid",
			entityType:     "agency",
			options:        []string{"valid"},
			row:            1,
			expectedErrors: 0,
			createFunc: func(entityType string, options []string, row int) *types.Gtfs {
				return CreateTestGtfsWithValidOptions(entityType, options, row)
			},
		},
		{
			name:           "Invalid",
			entityType:     "agency",
			options:        []string{"invalid"},
			row:            1,
			expectedErrors: 1,
			createFunc: func(entityType string, options []string, row int) *types.Gtfs {
				return CreateTestGtfsWithInvalidOptions(entityType, options, row)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gtfs := tt.createFunc(tt.entityType, tt.options, tt.row)
			verifyMapStructure(t, gtfs, tt.entityType, strings.Join(tt.options, ","), []int{tt.row}, tt.name)
		})
	}
}
