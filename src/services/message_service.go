package services

type Severity string

const (
	ErrorSeverity   Severity = "error"
	InfoSeverity    Severity = "info"
	WarningSeverity Severity = "warning"
)

type Message struct {
	Field    string   `json:"field"`
	FileName string   `json:"fileName"`
	Message  string   `json:"message"`
	Row      int      `json:"row"`
	RuleID   string   `json:"rule_id"`
	Severity Severity `json:"severity"`
}

type MessageService struct {
	errorCount   int
	infoCount    int
	warningCount int
	messages     []Message
}

func NewMessageService() *MessageService {
	return &MessageService{
		messages: []Message{},
	}
}

func (ms *MessageService) AddMessage(message Message) {
	ms.messages = append(ms.messages, message)

	switch message.Severity {
	case ErrorSeverity:
		ms.errorCount++
	case WarningSeverity:
		ms.warningCount++
	case InfoSeverity:
		ms.infoCount++
	}
}

type Summary struct {
	Messages      []Message `json:"messages"`
	TotalErrors   int       `json:"totalErrors"`
	TotalInfos    int       `json:"totalInfos"`
	TotalWarnings int       `json:"totalWarnings"`
}

func (ms *MessageService) GetSummary() Summary {
	return Summary{
		Messages:      ms.messages,
		TotalErrors:   ms.errorCount,
		TotalInfos:    ms.infoCount,
		TotalWarnings: ms.warningCount,
	}
}
