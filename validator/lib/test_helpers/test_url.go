package test_helpers

type UrlTestCase struct {
	Name           string
	Value          string
	Row            int
	ExpectedErrors int
}

func GetUrlTestCases() []UrlTestCase {
	return []UrlTestCase{
		{
			Name:           "Valid",
			Value:          "https://example.com",
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Invalid",
			Value:          "invalid-url",
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Missing Protocol",
			Value:          "example.com",
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Missing Host",
			Value:          "https://",
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Missing Path",
			Value:          "https://example.com",
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid Characters",
			Value:          "https://example.com/invalid#$%^&*()characters",
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:           "hasSpace",
			Value:          "https://example.com/with space",
			Row:            1,
			ExpectedErrors: 1,
		},
	}
}
