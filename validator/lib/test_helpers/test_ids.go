package test_helpers

import "main/lib"

// IdTestCase represents a test case for ID validation
type IdTestCase struct {
	Name           string
	Value          *string
	Row            int
	ExpectedErrors int
}

// GetIdTestCases returns all ID test cases with expected errors already set
func GetIdTestCases() []IdTestCase {

	return []IdTestCase{
		{
			Name:           "EmptyRequired",
			Value:          lib.Ptr(""),
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "NilRequired",
			Value:          nil,
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Unique",
			Value:          lib.Ptr("unique"),
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Duplicate",
			Value:          lib.Ptr("duplicate"),
			Row:            2,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid",
			Value:          lib.Ptr("invalid"),
			Row:            1,
			ExpectedErrors: 1,
		},
	}
}
