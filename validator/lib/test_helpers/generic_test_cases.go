package test_helpers

import (
	"main/lib"
	"main/types"
)

// ===============================================
// GENERIC TEST CASE STRUCTURES
// ===============================================

// GenericFieldTestCase is a universal test case for single field validation
type GenericFieldTestCase struct {
	Name             string
	Value            interface{} // Can be *string, *int, or any pointer type
	Row              int
	Rules            interface{} // Optional rules configuration
	ExpectedErrors   int
	ExpectedWarnings int
	ExpectedCodes    []string
}

// GenericRequiredFieldTestCase for testing required/recommended field patterns
type GenericRequiredFieldTestCase struct {
	Name             string
	Value            *string
	Row              int
	IsRequired       bool // Row 1 = required, Row 2+ = recommended (by convention)
	ExpectedErrors   int
	ExpectedWarnings int
	ExpectedCode     string
}

// GenericUrlTestCase for URL validation patterns
type GenericUrlTestCase struct {
	Name             string
	Url              *string
	Row              int
	ExpectedErrors   int
	ExpectedWarnings int
	ExpectedCode     string
}

// GenericColorTestCase for color field validation (6 char hex)
type GenericColorTestCase struct {
	Name             string
	Color            *string
	Row              int
	Severity         string // "error", "warning"
	ExpectedErrors   int
	ExpectedWarnings int
}

// GenericEnumTestCase for enum/option field validation
type GenericEnumTestCase struct {
	Name             string
	Value            interface{} // Can be *int or *string
	ValidOptions     []int
	Row              int
	ExpectedErrors   int
	ExpectedWarnings int
}

// GenericIdTestCase for ID field validation (required, unique)
type GenericIdTestCase struct {
	Name           string
	Id             *string
	Row            int
	ExistingIds    map[string][]int // For duplicate detection
	ExpectedErrors int
	ExpectedCode   string
}

// GenericSeverityTestCase for testing severity configurations
type GenericSeverityTestCase struct {
	Name             string
	Value            interface{}
	Severity         types.Severity // types.SEVERITY_ERROR, SEVERITY_WARNING, SEVERITY_IGNORE, SEVERITY_FORBIDDEN
	Row              int
	ExpectedErrors   int
	ExpectedWarnings int
}

// ===============================================
// GENERIC TEST CASE GENERATORS
// ===============================================

// GetGenericRequiredFieldTestCases returns test cases for required/recommended field patterns
func GetGenericRequiredFieldTestCases(fieldName string) []GenericRequiredFieldTestCase {
	return []GenericRequiredFieldTestCase{
		{
			Name:           "Required",
			Value:          nil,
			Row:            1,
			IsRequired:     true,
			ExpectedErrors: 1,
			ExpectedCode:   fieldName + "_required",
		},
		{
			Name:             "Recommended_Missing",
			Value:            nil,
			Row:              2,
			IsRequired:       false,
			ExpectedWarnings: 1,
			ExpectedCode:     fieldName + "_recommended",
		},
		{
			Name:           "Valid_Present",
			Value:          lib.Ptr("valid_value"),
			Row:            1,
			IsRequired:     true,
			ExpectedErrors: 0,
		},
	}
}

// GetGenericUrlTestCases returns comprehensive URL validation test cases
func GetGenericUrlTestCases(fieldName string) []GenericUrlTestCase {
	baseRow := 1

	cases := []GenericUrlTestCase{
		{
			Name:           "Valid_Url_Https",
			Url:            lib.Ptr("https://example.com"),
			Row:            baseRow,
			ExpectedErrors: 0,
		},
		{
			Name:           "Valid_Url_Http",
			Url:            lib.Ptr("http://example.com"),
			Row:            baseRow,
			ExpectedErrors: 0,
		},
		{
			Name:           "Valid_Url_With_Path",
			Url:            lib.Ptr("https://example.com/path/to/page"),
			Row:            baseRow,
			ExpectedErrors: 0,
		},
		{
			Name:           "Valid_Url_With_Query",
			Url:            lib.Ptr("https://example.com/page?param=value"),
			Row:            baseRow,
			ExpectedErrors: 0,
		},
		{
			Name:           "Invalid_Url_No_Protocol",
			Url:            lib.Ptr("example.com"),
			Row:            baseRow,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Url_No_Host",
			Url:            lib.Ptr("https://"),
			Row:            baseRow,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Url_Malformed",
			Url:            lib.Ptr("not-a-url"),
			Row:            baseRow,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Url_With_Space",
			Url:            lib.Ptr("https://example.com/path with space"),
			Row:            baseRow,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Url_Special_Characters",
			Url:            lib.Ptr("https://example.com/<script>"),
			Row:            baseRow,
			ExpectedErrors: 1,
		},
	}

	// Add required/missing cases
	cases = append([]GenericUrlTestCase{
		{
			Name:           "Required_Missing",
			Url:            nil,
			Row:            1,
			ExpectedErrors: 1,
			ExpectedCode:   fieldName + "_required",
		},
	}, cases...)

	return cases
}

// GetGenericColorTestCases returns test cases for hex color validation
func GetGenericColorTestCases(fieldName string) []GenericColorTestCase {
	return []GenericColorTestCase{
		{
			Name:           "Valid_Color_Uppercase",
			Color:          lib.Ptr("FFFFFF"),
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Valid_Color_Lowercase",
			Color:          lib.Ptr("ffffff"),
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Valid_Color_Mixed",
			Color:          lib.Ptr("123AbC"),
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Valid_Color_Black",
			Color:          lib.Ptr("000000"),
			Row:            1,
			ExpectedErrors: 0,
		},
		{
			Name:           "Invalid_Color_Too_Short",
			Color:          lib.Ptr("FFF"),
			Row:            2,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Color_Too_Long",
			Color:          lib.Ptr("FFFFFFF"),
			Row:            3,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Color_Bad_Characters",
			Color:          lib.Ptr("ZZZZZZ"),
			Row:            4,
			ExpectedErrors: 1,
		},
		{
			Name:           "Invalid_Color_With_Hash",
			Color:          lib.Ptr("#FFFFFF"),
			Row:            5,
			ExpectedErrors: 1,
		},
		{
			Name:           "Nil_Color_Optional",
			Color:          nil,
			Row:            7,
			ExpectedErrors: 0,
		},
	}
}

// GetGenericEnumIntTestCases returns test cases for integer enum validation
func GetGenericEnumIntTestCases(fieldName string, validOptions []int) []GenericEnumTestCase {
	cases := []GenericEnumTestCase{
		{
			Name:           "Missing_Value_Required",
			Value:          (*int)(nil),
			ValidOptions:   validOptions,
			Row:            1,
			ExpectedErrors: 1,
		},
	}

	// Add valid cases for each option
	for i, opt := range validOptions {
		val := opt
		cases = append(cases, GenericEnumTestCase{
			Name:           "Valid_Option_" + string(rune('0'+opt)),
			Value:          &val,
			ValidOptions:   validOptions,
			Row:            i + 2,
			ExpectedErrors: 0,
		})
	}

	// Add invalid case
	invalidVal := 999
	cases = append(cases, GenericEnumTestCase{
		Name:           "Invalid_Option",
		Value:          &invalidVal,
		ValidOptions:   validOptions,
		Row:            len(validOptions),
		ExpectedErrors: 1,
	})

	return cases
}

// GetGenericIdTestCases returns test cases for ID validation
func GetGenericIdTestCases(fieldName string) []GenericIdTestCase {
	return []GenericIdTestCase{
		{
			Name:           "Required Id",
			Id:             nil,
			Row:            1,
			ExistingIds:    nil,
			ExpectedErrors: 1,
			ExpectedCode:   fieldName + "_id_required",
		},
		{
			Name:           "Valid_Unique",
			Id:             lib.Ptr("unique_id"),
			Row:            1,
			ExistingIds:    map[string][]int{"unique_id": {0}},
			ExpectedErrors: 0,
		},
		{
			Name:           "Duplicate_Id",
			Id:             lib.Ptr("dup_id"),
			Row:            2,
			ExistingIds:    map[string][]int{"dup_id": {0, 1}},
			ExpectedErrors: 1,
		},
	}
}

// GetGenericSeverityTestCases returns test cases for different severity levels
func GetGenericSeverityTestCases(fieldName string) []GenericSeverityTestCase {
	return []GenericSeverityTestCase{
		{
			Name:           "Severity_Error_Missing",
			Value:          (*string)(nil),
			Severity:       types.SEVERITY_ERROR,
			Row:            1,
			ExpectedErrors: 1,
		},
		{
			Name:             "Severity_Warning_Missing",
			Value:            (*string)(nil),
			Severity:         types.SEVERITY_WARNING,
			Row:              2,
			ExpectedWarnings: 1,
		},
		{
			Name:             "Severity_Ignore_Missing",
			Value:            (*string)(nil),
			Severity:         types.SEVERITY_IGNORE,
			Row:              3,
			ExpectedErrors:   0,
			ExpectedWarnings: 0,
		},
		{
			Name:           "Severity_Forbidden_Present",
			Value:          lib.Ptr("present_value"),
			Severity:       types.SEVERITY_FORBIDDEN,
			Row:            4,
			ExpectedErrors: 1,
		},
		{
			Name:             "Severity_Forbidden_Missing",
			Value:            (*string)(nil),
			Severity:         types.SEVERITY_FORBIDDEN,
			Row:              5,
			ExpectedErrors:   0,
			ExpectedWarnings: 0,
		},
	}
}

// ===============================================
// GTFS-SPECIFIC COMMON TEST CASES
// ===============================================

// GetRouteTypeValidOptions returns valid GTFS route_type values
func GetRouteTypeValidOptions() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 11, 12}
}

// GetContinuousPickupDropOffValidOptions returns valid continuous_pickup/drop_off values
func GetContinuousPickupDropOffValidOptions() []int {
	return []int{0, 1, 2, 3}
}

// GetPickupTypeValidOptions returns valid pickup_type values for stop_times
func GetPickupTypeValidOptions() []int {
	return []int{0, 1, 2, 3}
}

// GetDropOffTypeValidOptions returns valid drop_off_type values for stop_times
func GetDropOffTypeValidOptions() []int {
	return []int{0, 1, 2, 3}
}

// GetTimepointValidOptions returns valid timepoint values for stop_times
func GetTimepointValidOptions() []int {
	return []int{0, 1}
}

// GetLocationTypeValidOptions returns valid location_type values for stops
func GetLocationTypeValidOptions() []int {
	return []int{0, 1, 2, 3, 4}
}

// GetWheelchairBoardingValidOptions returns valid wheelchair_boarding values
func GetWheelchairBoardingValidOptions() []int {
	return []int{0, 1, 2}
}

// GetBikesAllowedValidOptions returns valid bikes_allowed values
func GetBikesAllowedValidOptions() []int {
	return []int{0, 1, 2}
}

// GetDirectionIdValidOptions returns valid direction_id values
func GetDirectionIdValidOptions() []int {
	return []int{0, 1}
}

// GetExceptionTypeValidOptions returns valid exception_type values for calendar_dates
func GetExceptionTypeValidOptions() []int {
	return []int{1, 2}
}

// GetShapeDistTraveledValidOptions returns valid shape_dist_traveled values
func GetShapeDistTraveledValidOptions() []float64 {
	return []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
}

// GetShapePtLatValidOptions returns valid shape_pt_lat values
func GetShapePtLatValidOptions() []float32 {
	return []float32{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
}

// GetShapePtLngValidOptions returns valid shape_pt_lon values
func GetShapePtLonValidOptions() []float32 {
	return []float32{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
}

// GetDateValidOptions returns valid date values
func GetDateValidOptions() []string {
	return []string{
		"20240101",
		"20240102",
		"20240103",
	}
}

// ===============================================
// COMMON TEST VALUE GENERATORS
// ===============================================

// ValidTimezones returns a list of valid IANA timezones for testing
func GetValidTimezones() []string {
	return []string{
		"America/New_York",
		"America/Los_Angeles",
		"America/Chicago",
		"Europe/London",
		"Europe/Paris",
		"Europe/Lisbon",
		"Asia/Tokyo",
		"Asia/Shanghai",
		"Australia/Sydney",
		"Pacific/Auckland",
		"UTC",
	}
}

// GetInvalidTimezones returns invalid timezone strings for testing
func GetInvalidTimezones() []string {
	return []string{
		"Invalid/Timezone",
		"Not_A_Timezone",
		"America",
		"UTC+5",
		"EST",
		"PST",
	}
}

// GetValidLanguageCodes returns valid BCP-47 language codes
func GetValidLanguageCodes() []string {
	return []string{
		"en",
		"en-US",
		"en-GB",
		"pt",
		"pt-PT",
		"pt-BR",
		"es",
		"fr",
		"de",
		"zh",
		"ja",
	}
}

// GetInvalidLanguageCodes returns invalid language codes
func GetInvalidLanguageCodes() []string {
	return []string{
		"invalid-lang",
		"english",
		"123",
		"en_US",
		"",
	}
}

// GetValidEmails returns valid email addresses for testing
func GetValidEmails() []string {
	return []string{
		"test@example.com",
		"user.name@domain.org",
		"contact+tag@company.co.uk",
		"admin@subdomain.example.com",
	}
}

// GetInvalidEmails returns invalid email addresses for testing
func GetInvalidEmails() []string {
	return []string{
		"invalid-email",
		"@domain.com",
		"user@",
		"user@.com",
		"user name@domain.com",
		"user@domain",
	}
}

// GetValidPhoneNumbers returns sample phone numbers for testing
func GetValidPhoneNumbers() []string {
	return []string{
		"123-456-7890",
		"+1-234-567-8901",
		"(555) 123-4567",
		"+44 20 7946 0958",
	}
}

// GetValidUrls returns sample valid URLs for testing
func GetValidUrls() []string {
	return []string{
		"https://example.com",
		"http://example.org",
		"https://www.example.com/path",
		"https://example.com/page?query=value",
		"https://subdomain.example.com",
	}
}

// GetInvalidUrls returns sample invalid URLs for testing
func GetInvalidUrls() []string {
	return []string{
		"invalid-url",
		"not a url",
		"ftp://example.com",
		"//example.com",
		"example.com",
		"https://",
		"https://example.com/path with space",
	}
}

// ===============================================
// TEST CASE COUNT SUMMARY
// ===============================================

// TestCaseSummary provides counts of all available test cases
type TestCaseSummary struct {
	RequiredFieldCases int
	UrlCases           int
	ColorCases         int
	EnumCases          int
	IdCases            int
	SeverityCases      int
	TotalCases         int
}

// GetTestCaseSummary returns a summary of all available generic test cases
func GetTestCaseSummary() TestCaseSummary {
	requiredCases := len(GetGenericRequiredFieldTestCases("sample"))
	urlCases := len(GetGenericUrlTestCases("sample"))
	colorCases := len(GetGenericColorTestCases("sample"))
	enumCases := len(GetGenericEnumIntTestCases("sample", []int{0, 1, 2}))
	idCases := len(GetGenericIdTestCases("sample"))
	severityCases := len(GetGenericSeverityTestCases("sample"))

	return TestCaseSummary{
		RequiredFieldCases: requiredCases,
		UrlCases:           urlCases,
		ColorCases:         colorCases,
		EnumCases:          enumCases,
		IdCases:            idCases,
		SeverityCases:      severityCases,
		TotalCases:         requiredCases + urlCases + colorCases + enumCases + idCases + severityCases,
	}
}
