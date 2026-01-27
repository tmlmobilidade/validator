package test_helpers

import (
	"main/lib"
	"main/types"
	"testing"
)

type UrlTestCase struct {
	Name           string
	Gtfs           *types.Gtfs
	Agency         *types.Agency
	Row            int
	ExpectedErrors int
}

func GetUrlTestCases() []UrlTestCase {
	return []UrlTestCase{
		{
			Name:           "Valid",
			Gtfs:           CreateTestGtfsValidUrl("agency", "https://example.com", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("https://example.com")},
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Invalid",
			Gtfs:           CreateTestGtfsInvalidUrl("agency", "invalid-url", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("invalid-url")},
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Missing Protocol",
			Gtfs:           CreateTestGtfsUrlWithMissingProtocol("agency", "example.com", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("example.com")},
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Missing Host",
			Gtfs:           CreateTestGtfsUrlWithMissingHost("agency", "https://", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("https://")},
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Missing Path",
			Gtfs:           CreateTestGtfsUrlWithMissingPath("agency", "https://example.com", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("https://example.com")},
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid Characters",
			Gtfs:           CreateTestGtfsUrlWithInvalidCharacters("agency", "https://example.com/invalid#$%^&*()characters", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("https://example.com/invalid#$%^&*()characters")},
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "hasSpace",
			Gtfs:           CreateTestGtfsUrlWithSpace("agency", "https://example.com/with space", 1),
			Agency:         &types.Agency{AgencyUrl: lib.Ptr("https://example.com/with space")},
			Row:            1,
			ExpectedErrors: 1,
		},
	}
}
func CreateTestGtfsValidUrl(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func CreateTestGtfsInvalidUrl(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func CreateTestGtfsUrlWithSpace(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func CreateTestGtfsUrlWithInvalidCharacters(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func CreateTestGtfsUrlWithMissingProtocol(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func CreateTestGtfsUrlWithMissingHost(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func CreateTestGtfsUrlWithMissingPath(entityType string, url string, row int) *types.Gtfs {
	return &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"agency": {
				url: {1},
			},
		},
	}
}

func TestAllUrlHelpers(t *testing.T) {
	t.Helper()

	tests := []struct {
		name           string
		entityType     string
		url            string
		row            []int
		createFunc     func(string, string, ...int) *types.Gtfs
		expectedErrors int
	}{
		{
			name:       "Valid",
			entityType: "agency",
			url:        "https://example.com",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsValidUrl(entityType, url, rows[0])
			},
			expectedErrors: 0,
		},
		{
			name:       "Invalid",
			entityType: "agency",
			url:        "invalid-url",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsInvalidUrl(entityType, url, rows[0])
			},
			expectedErrors: 1,
		},
		{
			name:       "Missing Protocol",
			entityType: "agency",
			url:        "example.com",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsUrlWithMissingProtocol(entityType, url, rows[0])
			},
			expectedErrors: 1,
		},
		{
			name:       "Missing Host",
			entityType: "agency",
			url:        "https://",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsUrlWithMissingHost(entityType, url, rows[0])
			},
			expectedErrors: 1,
		},
		{
			name:       "Missing Path",
			entityType: "agency",
			url:        "https://example.com",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsUrlWithMissingPath(entityType, url, rows[0])
			},
			expectedErrors: 1,
		},
		{
			name:       "Invalid Characters",
			entityType: "agency",
			url:        "https://example.com/invalid#$%^&*()characters",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsUrlWithInvalidCharacters(entityType, url, rows[0])
			},
			expectedErrors: 1,
		},
		{
			name:       "hasSpace",
			entityType: "agency",
			url:        "https://example.com/with space",
			row:        []int{1},
			createFunc: func(entityType string, url string, rows ...int) *types.Gtfs {
				return CreateTestGtfsUrlWithSpace(entityType, url, rows[0])
			},
			expectedErrors: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gtfs := tt.createFunc(tt.entityType, tt.url, tt.row...)
			verifyMapStructure(t, gtfs, tt.entityType, tt.url, tt.row, tt.name)
		})
	}
}
