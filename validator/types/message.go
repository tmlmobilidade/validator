package types

type Severity string

const (
	SEVERITY_ERROR   Severity = "error"
	SEVERITY_INFO    Severity = "info"
	SEVERITY_WARNING Severity = "warning"
)

type Message struct {
	Row          int      `json:"row"`
	Field        string   `json:"field"`
	FileName     string   `json:"fileName"`
	Message      string   `json:"message"`
	ValidationID string   `json:"validation_id"`
	Severity     Severity `json:"severity"`
}

type Summary struct {
	Messages      []Message `json:"messages"`
	TotalErrors   int       `json:"totalErrors"`
	TotalInfos    int       `json:"totalInfos"`
	TotalWarnings int       `json:"totalWarnings"`
}