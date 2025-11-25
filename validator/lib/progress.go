package lib

import (
	"fmt"
	"main/types"
)

// ProgressTracker handles progress tracking and logging for validation iterations
type ProgressTracker struct {
	tableName         string
	totalCount        int
	processedCount    int
	lastLoggedPercent int
	rowThreshold      int // Log every N rows when totalCount is unknown
}

// NewProgressTracker creates a new progress tracker for a table
func NewProgressTracker(tableName string, totalCount int, rowThreshold int) *ProgressTracker {
	return &ProgressTracker{
		tableName:         tableName,
		totalCount:        totalCount,
		processedCount:    0,
		lastLoggedPercent: -1,
		rowThreshold:      rowThreshold,
	}
}

// Track increments the processed count and logs progress if needed
func (pt *ProgressTracker) Track() {
	pt.processedCount++

	// Log progress every 10% or every rowThreshold rows (whichever comes first)
	if pt.totalCount > 0 {
		currentPercent := (pt.processedCount * 100) / pt.totalCount
		if currentPercent != pt.lastLoggedPercent && (currentPercent%10 == 0 || pt.processedCount%pt.rowThreshold == 0) {
			AppLogger.Debug(fmt.Sprintf("Validating %s: %d/%d (%.1f%%)", pt.tableName, pt.processedCount, pt.totalCount, float64(pt.processedCount)*100.0/float64(pt.totalCount)))
			pt.lastLoggedPercent = currentPercent
		}
	} else if pt.processedCount%pt.rowThreshold == 0 {
		AppLogger.Debug(fmt.Sprintf("Validating %s: %d rows processed", pt.tableName, pt.processedCount))
	}
}

// GetProcessedCount returns the number of processed rows
func (pt *ProgressTracker) GetProcessedCount() int {
	return pt.processedCount
}

// GetTotalCount returns the total count
func (pt *ProgressTracker) GetTotalCount() int {
	return pt.totalCount
}

// CreateProgressTracker is a convenience function that creates a progress tracker
// by getting the table count from gtfs
func CreateProgressTracker(gtfs types.Gtfs, tableName string, rowThreshold int) *ProgressTracker {
	totalCount, err := gtfs.GetTableCount(tableName)
	if err != nil {
		AppLogger.Debug(fmt.Sprintf("Could not get table count for %s: %v", tableName, err))
		totalCount = 0
	}
	return NewProgressTracker(tableName, totalCount, rowThreshold)
}
