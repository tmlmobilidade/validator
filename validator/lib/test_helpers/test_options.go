package test_helpers

type OptionTestCase struct {
	Name           string
	Value          int
	AllowedOptions []int
	ExpectedErrors int
}

func GetOptionTestCases(allowedOptions []int) []OptionTestCase {
	validValue := allowedOptions[0]
	invalidValue := allowedOptions[0] + 100

	return []OptionTestCase{
		{
			Name:           "Valid",
			Value:          validValue,
			AllowedOptions: allowedOptions,
			ExpectedErrors: 0,
		},
		{
			Name:           "Invalid",
			Value:          invalidValue,
			AllowedOptions: allowedOptions,
			ExpectedErrors: 1,
		},
	}
}
