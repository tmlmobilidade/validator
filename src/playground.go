// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"main/src/lib"
// 	"main/src/types"
// )

// type Agency struct {
// 	AgencyEmail *string `json:"agency_email,omitempty"`
// 	AgencyFareUrl *string `json:"agency_fare_url,omitempty"`
// 	AgencyId string `json:"agency_id"`
// 	AgencyLang *string `json:"agency_lang,omitempty"`
// 	AgencyName string `json:"agency_name"`
// 	AgencyPhone *string `json:"agency_phone,omitempty"`
// 	AgencyTimezone string `json:"agency_timezone"`
// 	AgencyUrl string `json:"agency_url"`
// }

// // type Agency struct {
// // 	AgencyEmail *string `json:"agency_email"`
// // 	AgencyFareUrl *string `json:"agency_fare_url"`
// // 	AgencyId string `json:"agency_id"`
// // 	AgencyLang *string `json:"agency_lang"`
// // 	AgencyName string `json:"agency_name"`
// // 	AgencyPhone *string `json:"agency_phone"`
// // 	AgencyTimezone string `json:"agency_timezone"`
// // 	AgencyUrl string `json:"agency_url"`
// // }

// func CreateAgencyFromMap(m map[string]string) (a Agency, errors []string) {

// 	errors = []string{}
// 	item := Agency{}

// 	//Convert Optional Values
// 	var agencyEmail, agencyFareUrl, agencyLang, agencyPhone string
// 	lib.ParseStringToPrimitive(m["agency_email"], &agencyEmail, &errors)
// 	lib.ParseStringToPrimitive(m["agency_fare_url"], &agencyFareUrl, &errors)
// 	lib.ParseStringToPrimitive(m["agency_lang"], &agencyLang, &errors)
// 	lib.ParseStringToPrimitive(m["agency_phone"], &agencyPhone, &errors)

// 	item.AgencyEmail = &agencyEmail
// 	item.AgencyFareUrl = &agencyFareUrl
// 	item.AgencyLang = &agencyLang
// 	item.AgencyPhone = &agencyPhone

// 	//Convert Required Values
// 	lib.ParseStringToPrimitive(m["agency_timezone"], &item.AgencyTimezone, &errors)
// 	lib.ParseStringToPrimitive(m["agency_name"], &item.AgencyName, &errors)
// 	lib.ParseStringToPrimitive(m["agency_id"], &item.AgencyId, &errors)
// 	lib.ParseStringToPrimitive(m["agency_url"], &item.AgencyUrl, &errors)

// 	// Validate Required Values
// 	requiredValues := map[string]string{
// 		"agency_timezone": item.AgencyTimezone,
// 		"agency_name": item.AgencyName,
// 		"agency_id": item.AgencyId,
// 		"agency_url": item.AgencyUrl,
// 	}

// 	for key, value := range requiredValues {
// 		if value == "" {
// 			errors = append(errors, fmt.Sprintf("Required value \"%s\" is empty", key))
// 		}
// 	}

// 	return item, errors
// }

// type Stop struct {
// 	LevelId* string `json:"level_id,omitempty"`
// 	LocationType* types.LocationType `json:"location_type,omitempty"`
// 	ParentStation* string `json:"parent_station,omitempty"`
// 	PlatformCode* string `json:"platform_code,omitempty"`
// 	StopCode* string `json:"stop_code,omitempty"`
// 	StopDesc* string `json:"stop_desc,omitempty"`
// 	StopId string `json:"stop_id"`
// 	StopLat* float32 `json:"stop_lat,omitempty"`
// 	StopLon* float32 `json:"stop_lon,omitempty"`
// 	StopName* string `json:"stop_name,omitempty"`
// 	StopTimezone* string `json:"stop_timezone,omitempty"`
// 	StopUrl* string `json:"stop_url,omitempty"`
// 	WheelchairBoarding* types.WheelchairBoardingType `json:"wheelchair_boarding,omitempty"`
// 	ZoneId* string `json:"zone_id,omitempty"`
// }

// func CreateStopFromMap(m map[string]string) (item Stop, errors []string) {

// 	errors = []string{}

// 	//Convert Optional Values
// 	var levelId, parentStation, platformCode, stopCode, stopDesc, stopName, stopTimezone, stopUrl, zoneId string

// 	lib.ParseStringToPrimitive(m["level_id"], &levelId, &errors)
// 	lib.ParseStringToPrimitive(m["parent_station"], &parentStation, &errors)
// 	lib.ParseStringToPrimitive(m["platform_code"], &platformCode, &errors)
// 	lib.ParseStringToPrimitive(m["stop_code"], &stopCode, &errors)
// 	lib.ParseStringToPrimitive(m["stop_desc"], &stopDesc, &errors)
// 	lib.ParseStringToPrimitive(m["stop_name"], &stopName, &errors)
// 	lib.ParseStringToPrimitive(m["stop_timezone"], &stopTimezone, &errors)
// 	lib.ParseStringToPrimitive(m["stop_url"], &stopUrl, &errors)
// 	lib.ParseStringToPrimitive(m["zone_id"], &zoneId, &errors)

// 	//Convert Optional Enums
// 	// lib.ParseStringToPrimitive(m["wheelchair_boarding"], &wheelchairBoarding, &errors)
// 	// if wheelchairBoarding != "" {
// 	// 	var wheelchairBoardingType types.WheelchairBoardingType
// 	// 	lib.ParseStringToPrimitive(wheelchairBoarding, &wheelchairBoardingType, &errors)
// 	// 	item.WheelchairBoarding = &wheelchairBoardingType
// 	// }

// 	item.LevelId = &levelId
// 	item.ParentStation = &parentStation
// 	item.PlatformCode = &platformCode
// 	item.StopCode = &stopCode
// 	item.StopDesc = &stopDesc
// 	item.StopName = &stopName
// 	item.StopTimezone = &stopTimezone
// 	item.StopUrl = &stopUrl
// 	item.ZoneId = &zoneId

// 	return item, errors
// }

// func main() {
// 	agencyMap := map[string]string{
// 		"agency_email": "test@test.com",
// 		"agency_fare_url": "https://test.com",
// 		"agency_lang": "en",
// 		"agency_timezone": "America/New_York",
// 		// "agency_url": "https://test.com",
// 		"agency_name": "Test Agency",
// 	}

// 	// stopMap := map[string]string{
// 	// 	"stop_lon": "123.456",
// 	// 	"stop_lat": "3243434343434",
// 	// 	"wheelchair_boarding": "1",
// 	// }

// 	errors := []string{}

// 	agency, err := CreateAgencyFromMap(agencyMap)
// 	errors = append(errors, err...)
// 	printMap(agency)

// 	// stop, err := CreateStopFromMap(stopMap)
// 	// errors = append(errors, err...)
// 	// printMap(stop)

// 	fmt.Println("errors:", errors)

// }

// func printMap(a any) {
// 	b, err := json.MarshalIndent(a, "", "  ")
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Printf("%s\n", string(b))
// }
