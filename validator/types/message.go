package types

type Severity string

const (
	SEVERITY_IGNORE  Severity = "ignore"
	SEVERITY_ERROR   Severity = "error"
	SEVERITY_INFO    Severity = "info"
	SEVERITY_WARNING Severity = "warning"
)

type Message struct {
	Rows         []int      `json:"rows"`
	Field        string     `json:"field"`
	FileName     string     `json:"file_name"`
	Message      string     `json:"message"`
	ValidationID string     `json:"validation_id"`
	Severity     Severity   `json:"severity"`
}

type Summary struct {
	Messages      []Message `json:"messages"`
	TotalErrors   int       `json:"total_errors"`
	TotalInfos    int       `json:"total_infos"`
	TotalWarnings int       `json:"total_warnings"`
}