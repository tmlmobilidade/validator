package services

import (
	"encoding/json"
	"fmt"
	"main/validator/types"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
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

func (ms *MessageService) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Validation ID", "Message", "Severity", "Field", "File Name", "Row"})
	table.SetRowSeparator("-")
	table.SetFooter([]string{"", "", "Errors: " + strconv.Itoa(ms.errorCount), "Infos: " + strconv.Itoa(ms.infoCount), "Warnings: " + strconv.Itoa(ms.warningCount), "Total: " + strconv.Itoa(ms.errorCount+ms.infoCount+ms.warningCount)})

	for _, message := range ms.messages {
		table.Append([]string{message.ValidationID, message.Message, string(message.Severity), message.Field, message.FileName, strconv.Itoa(message.Row)})
	}
	table.Render()
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

// PrintJSON writes the messages to a JSON file. Defaults to "messages.json" if no path is provided.
func (ms *MessageService) PrintJSON(outputPath ...string) {
	outputPathString := "messages.json"
	if len(outputPath) > 0 && outputPath[0] != "" {
		outputPathString = outputPath[0]
	}

	if len(ms.messages) == 0 {
		fmt.Println("No messages to print")
		return
	}

	data := struct {
		Messages []types.Message `json:"messages"`
	}{
		Messages: ms.messages,
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %v\n", err)
		return
	}

	if err := os.WriteFile(outputPathString, jsonBytes, 0644); err != nil {
		fmt.Printf("Failed to write JSON file: %v\n", err)
	}
}

var AppMessageService = NewMessageService()
