package services

import (
	"fmt"
	"main/src/types"
)

type MessageService struct {
	errorCount   int
	infoCount    int
	warningCount int
	messages     []types.Message
}

func NewMessageService() *MessageService {
	return &MessageService{
		messages: []types.Message{},
	}
}

func (ms *MessageService) AddMessage(message types.Message) {
	ms.messages = append(ms.messages, message)

	switch message.Severity {
	case types.SEVERITY_ERROR:
		ms.errorCount++
	case types.SEVERITY_WARNING:
		ms.warningCount++
	case types.SEVERITY_INFO:
		ms.infoCount++
	}
}

func (ms *MessageService) GetSummary() types.Summary {
	return types.Summary{
		Messages:      ms.messages,
		TotalErrors:   ms.errorCount,
		TotalInfos:    ms.infoCount,
		TotalWarnings: ms.warningCount,
	}
}

func (ms *MessageService) PrintSummary() {
	summary := ms.GetSummary()
	fmt.Println("\n\n================================================")
	fmt.Println("GTFS Validation Summary")
	fmt.Println("================================================")
	fmt.Printf("Total Errors: %d\n", summary.TotalErrors)
	fmt.Printf("Total Infos: %d\n", summary.TotalInfos)
	fmt.Printf("Total Warnings: %d\n", summary.TotalWarnings)
	fmt.Println("================================================")
}
