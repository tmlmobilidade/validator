package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type RGB [3]int

// Predefined log levels with associated colors
var (
	ColorDebug = RGB{255, 128, 0}   // Orange
	ColorInfo  = RGB{100, 149, 237} // Cornflower Blue
	ColorError = RGB{220, 20, 60}   // Crimson
)

// Logger represents a customizable logger
type Logger struct {
	WithTimestamp bool
}

// NewLogger creates a new Logger instance
func NewLogger(withTimestamp bool) *Logger {
	return &Logger{WithTimestamp: withTimestamp}
}

// log prints a message with the specified RGB color
func (l *Logger) log(message string, clr RGB) {
	if l.WithTimestamp {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		message = fmt.Sprintf("[%s] %s", timestamp, message)
	}
	color.RGB(clr[0], clr[1], clr[2]).Println(message)
}

// Divider prints a nice divider, with optional centered text
func (l *Logger) Divider(message ...string) {
	const width = 50
	line := strings.Repeat("-", width)

	if len(message) == 0 {
		fmt.Printf("\n%s\n", line)
		return
	}

	msg := strings.TrimSpace(message[0])
	padding := (width - len(msg) - 2) / 2
	if padding < 0 {
		padding = 0
	}

	fmt.Printf("\n%s %s %s\n", strings.Repeat("-", padding), msg, strings.Repeat("-", padding))
}

// Debug logs a debug-level message
func (l *Logger) Debug(message string) {
	l.log(message, ColorDebug)
}

// Info logs an info-level message
func (l *Logger) Info(message string) {
	l.log(message, ColorInfo)
}

// Error logs an error-level message
func (l *Logger) Error(message string) {
	l.log(message, ColorError)
}

// Custom logs a message with a custom color
func (l *Logger) Custom(message string, clr RGB) {
	l.log(message, clr)
}
