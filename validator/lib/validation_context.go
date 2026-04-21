package lib

import (
	"main/i18n"
	"main/types"
)

// MessageAdder is an interface for adding validation messages
// This avoids import cycles by not depending on services.MessageService directly
type MessageAdder interface {
	AddMessage(message types.Message)
}

// ValidationContext encapsulates common validation patterns and reduces duplication
// across validation functions. It handles message creation, severity checking,
// and common validation logic.
type ValidationContext struct {
	// Field name being validated (e.g., "stop_code")
	Field string

	// File name (e.g., "stops.txt")
	FileName string

	// Row number (0-based index)
	Row int

	// Validation ID (e.g., "stop_code_validation")
	ValidationID string

	// Rule ID (e.g., "stop_id_rule")
	RuleID string

	// Severity from rules (defaults to SEVERITY_IGNORE)
	Severity types.Severity

	// Message adder for adding messages (avoids import cycle)
	MessageAdder MessageAdder
}

// NewValidationContext creates a new ValidationContext with default severity
func NewValidationContext(field, fileName, validationID, ruleID string, row int, messageAdder MessageAdder) *ValidationContext {
	return &ValidationContext{
		Field:        field,
		FileName:     fileName,
		Row:          row,
		ValidationID: validationID,
		RuleID:       ruleID,
		Severity:     types.SEVERITY_IGNORE,
		MessageAdder: messageAdder,
	}
}

// WithSeverity sets the severity from rules and returns the context for chaining
func (vc *ValidationContext) WithSeverity(severity types.Severity) *ValidationContext {
	if severity != "" {
		vc.Severity = severity
	}
	return vc
}

// ShouldIgnore returns true if validation should be skipped (IGNORE severity)
func (vc *ValidationContext) ShouldIgnore() bool {
	return vc.Severity == types.SEVERITY_IGNORE
}

// IsForbidden returns true if the field is forbidden (FORBIDDEN severity)
func (vc *ValidationContext) IsForbidden() bool {
	return vc.Severity == types.SEVERITY_FORBIDDEN
}

// ShouldSkip returns true if validation should be skipped (IGNORE or FORBIDDEN)
func (vc *ValidationContext) ShouldSkip() bool {
	return vc.ShouldIgnore() || vc.IsForbidden()
}

// AddError adds an error message
func (vc *ValidationContext) AddError(message string) {
	vc.AddMessage(message, types.SEVERITY_ERROR)
}

// AddWarning adds a warning message
func (vc *ValidationContext) AddWarning(message string) {
	vc.AddMessage(message, types.SEVERITY_WARNING)
}

// AddMessage adds a message with the specified severity
func (vc *ValidationContext) AddMessage(message string, severity types.Severity) {
	vc.MessageAdder.AddMessage(types.Message{
		Field:        vc.Field,
		FileName:     vc.FileName,
		Rows:         []int{vc.Row},
		Message:      message,
		Severity:     severity,
		ValidationID: vc.ValidationID,
		RuleID:       vc.RuleID,
	})
}

// AddMessageWithSeverity adds a message using the context's severity
func (vc *ValidationContext) AddMessageWithSeverity(message string) {
	vc.AddMessage(message, vc.Severity)
}

// GetTranslatedMessage gets a translated message by key
func (vc *ValidationContext) GetTranslatedMessage(key string, args ...interface{}) string {
	return i18n.AppTranslator.Get(key, args...)
}

// GetRequiredMessage gets a translated message for required field validation
func (vc *ValidationContext) GetRequiredMessage(requiredKey, recommendedKey string) string {
	if vc.Severity == types.SEVERITY_ERROR {
		return vc.GetTranslatedMessage(requiredKey)
	}
	return vc.GetTranslatedMessage(recommendedKey)
}
