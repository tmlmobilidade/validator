package types

type Validation struct {
	ID          string
	Description string
	Severity    Severity

	Validate func(gtfsData Gtfs) []Message
}
