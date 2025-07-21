package services

import (
	"encoding/json"
	"fmt"
	"main/lib"
	"main/types"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const TOTAL_ISSUES_LIMIT = 500

type MessageService struct {
	errorCount   int
	warningCount int
	messages     []types.Message
}

func NewMessageService() *MessageService {
	return &MessageService{
		messages: []types.Message{},
	}
}

func (ms *MessageService) AddMessages(messages []types.Message) {
	for _, message := range messages {
		ms.AddMessage(message)
	}
}

func (ms *MessageService) AddMessage(message types.Message) {

	// Add +2 to each row in the message.Rows
	// 1 for the header and 1 for the 0 based index
	for i, row := range message.Rows {
		message.Rows[i] = row + 2
	}

	for i, m := range ms.messages {
		if m.Message == message.Message {
			// Only keep up to 100 rows, keeping the latest row
			newRows := append(m.Rows, message.Rows...)
			if len(newRows) > 100 {
				// Keep first 99 rows and the latest row
				lastRow := newRows[len(newRows)-1]
				newRows = append(newRows[:99], lastRow)
			}
			ms.messages[i].Rows = newRows
			return
		}
	}

	ms.messages = append(ms.messages, message)

	switch message.Severity {
	case types.SEVERITY_ERROR:
		ms.errorCount++
	case types.SEVERITY_WARNING:
		ms.warningCount++
	}

	// Exit if total errors + warnings exceeds TOTAL_ISSUES_LIMIT
	if ms.errorCount+ms.warningCount >= TOTAL_ISSUES_LIMIT {
		lib.AppLogger.Error("Too many issues (errors + warnings > " + strconv.Itoa(TOTAL_ISSUES_LIMIT) + "). Exiting.")
		ms.PrintJSON()
		os.Exit(0)
	}
}

func (ms *MessageService) GetSummary() types.Summary {
	return types.Summary{
		Messages:      ms.messages,
		TotalErrors:   ms.errorCount,
		TotalWarnings: ms.warningCount,
	}
}

func (ms *MessageService) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Validation ID", "Message", "Severity", "Field", "File Name", "Row"})
	table.SetRowSeparator("-")
	table.SetFooter([]string{"", "", "Errors: " + strconv.Itoa(ms.errorCount), "Warnings: " + strconv.Itoa(ms.warningCount), "Total: " + strconv.Itoa(ms.errorCount+ms.warningCount), ""})
	for _, message := range ms.messages {
		rows := make([]string, len(message.Rows))
		for i, row := range message.Rows {
			rows[i] = strconv.Itoa(row)
		}
		table.Append([]string{message.ValidationID, message.Message, string(message.Severity), message.Field, message.FileName, strings.Join(rows, ", ")})
	}
	table.Render()
}

func (ms *MessageService) PrintSummary() {
	summary := ms.GetSummary()
	fmt.Println("\n\n================================================")
	fmt.Println("GTFS Validation Summary")
	fmt.Println("================================================")
	fmt.Printf("Total Errors: %d\n", summary.TotalErrors)
	fmt.Printf("Total Warnings: %d\n", summary.TotalWarnings)
	fmt.Println("================================================")
}

func (ms *MessageService) PrintJSON() {
	lib.PrintMap(ms.GetSummary(), true)
}

func (ms *MessageService) WriteToFile(filename string) {
	json, err := json.Marshal(ms.GetSummary())
	if err != nil {
		lib.AppLogger.Error("Error marshalling summary to JSON: " + err.Error())
		return
	}

	os.WriteFile(filename, json, 0644)
}

func (ms *MessageService) Clear() {
	ms.messages = []types.Message{}
	ms.errorCount = 0
	ms.warningCount = 0
}

var AppMessageService = NewMessageService()
