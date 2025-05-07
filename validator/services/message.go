package services

import (
	"fmt"
	"main/lib"
	"main/types"
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
	for i, m := range ms.messages {
		if m.ValidationID == message.ValidationID && m.Field == message.Field && m.FileName == message.FileName {
			// Add +2 to each row in the message.Rows
			for i, row := range message.Rows {
				message.Rows[i] = row + 2 // Add +2, 1 for the header and 1 for the 0 based index
			}
			
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
		table.Append([]string{message.ValidationID, message.Message, string(message.Severity), message.Field, message.FileName, strconv.Itoa(message.Rows[0])})
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

func (ms *MessageService) PrintJSON() {
	lib.PrintMap(ms.GetSummary(), true)
}

var AppMessageService = NewMessageService()
