package config

// Database configuration constants
const (
	// BatchSize is the number of rows to insert in a single transaction
	// Tune this depending on SSD/HDD performance
	BatchSize = 2000
)

// Validation configuration constants
const (
	// TotalIssuesLimit is the maximum number of errors + warnings before validation stops
	TotalIssuesLimit = 500
)

// Progress tracking configuration constants
const (
	// ProgressThresholdLarge is the row threshold for logging progress on large tables
	// Used for tables like stops, trips, routes, stop_times, shapes
	ProgressThresholdLarge = 1000

	// ProgressThresholdSmall is the row threshold for logging progress on small tables
	// Used for tables like agency, feed_info, calendar, calendar_dates, fare_rules, fare_attributes
	ProgressThresholdSmall = 100
)

