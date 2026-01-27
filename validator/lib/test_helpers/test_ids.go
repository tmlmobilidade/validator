package test_helpers

import (
	"main/lib"
	"main/types"
	"testing"
)

// IdTestCase represents a test case for ID validation
type IdTestCase struct {
	Name           string
	Gtfs           *types.Gtfs
	Agency         *types.Agency
	Row            int
	ExpectedErrors int
}

// GetIdTestCases returns all ID test cases with expected errors already set
func GetIdTestCases() []IdTestCase {
	return []IdTestCase{
		{
			Name:           "Unique",
			Gtfs:           CreateTestGtfsWithUniqueId("agency", "unique", 1),
			Agency:         &types.Agency{AgencyId: lib.Ptr("unique")},
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Duplicate",
			Gtfs:           CreateTestGtfsWithDuplicateId("agency", "duplicate", 1, 2),
			Agency:         &types.Agency{AgencyId: lib.Ptr("duplicate")},
			Row:            2,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid",
			Gtfs:           CreateTestGtfsInvalidId("agency", "invalid", 1),
			Agency:         &types.Agency{AgencyId: lib.Ptr("invalid")},
			Row:            1,
			ExpectedErrors: 1,
		},
	}
}

// CreateTestGtfsWithUniqueId creates a Gtfs instance with an IdMap configured for a unique ID
func CreateTestGtfsWithUniqueId(entityType string, id string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			entityType: {
				id: {row},
			},
		},
	}
}

// CreateTestGtfsWithDuplicateId creates a Gtfs instance with an IdMap configured for a duplicate ID
func CreateTestGtfsWithDuplicateId(entityType string, id string, rows ...int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			entityType: {
				id: rows,
			},
		},
	}
}

// CreateTestGtfsInvalidId creates a Gtfs instance with an IdMap configured for an invalid (duplicate) ID
func CreateTestGtfsInvalidId(entityType string, id string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			entityType: {
				id: {row, row},
			},
		},
	}
}

// TestAllIdHelpers tests all ID helper functions to ensure they work correctly
// Uses table-driven tests for better performance and maintainability
// This function validates:
// - CreateTestGtfsWithUniqueId
// - CreateTestGtfsWithDuplicateId
// - CreateTestGtfsValidId
// - CreateTestGtfsInvalidId

func TestAllIdHelpers(t *testing.T) {
	t.Helper()

	tests := []struct {
		name           string
		entityType     string
		id             string
		rows           []int
		expectedErrors int
		createFunc     func(string, string, ...int) *types.Gtfs
	}{
		{
			name:           "Unique",
			entityType:     "agency",
			id:             "unique",
			rows:           []int{1},
			expectedErrors: 0,
			createFunc: func(entityType string, id string, rows ...int) *types.Gtfs {
				return CreateTestGtfsWithUniqueId(entityType, id, rows[0])
			},
		},
		{
			name:           "Duplicate",
			entityType:     "agency",
			id:             "duplicate",
			rows:           []int{1, 2},
			expectedErrors: 1,
			createFunc: func(entityType string, id string, rows ...int) *types.Gtfs {
				return CreateTestGtfsWithDuplicateId(entityType, id, rows...)
			},
		},
		{
			name:           "Invalid",
			entityType:     "agency",
			id:             "invalid",
			rows:           []int{1},
			expectedErrors: 1,
			createFunc: func(entityType string, id string, rows ...int) *types.Gtfs {
				return CreateTestGtfsInvalidId(entityType, id, rows[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gtfs := tt.createFunc(tt.entityType, tt.id, tt.rows...)
			verifyMapStructure(t, gtfs, tt.entityType, tt.id, tt.rows, tt.name)
		})
	}
}
