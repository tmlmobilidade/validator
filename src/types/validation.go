package types

type Validation struct {
	ID          string
	Description string
	Severity    Severity

	Validate func(gtfs Gtfs) []Message
}
