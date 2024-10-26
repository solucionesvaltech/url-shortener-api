package customerror

import (
	"fmt"
)

type ErrorType string

const (
	Setup      ErrorType = "SETUP_ERROR"
	Mapping    ErrorType = "MAPPING_ERROR"
	Parsing    ErrorType = "PARSING_ERROR"
	Validation ErrorType = "VALIDATION_ERROR"
	Processing ErrorType = "PROCESSING_ERROR"
	Saving     ErrorType = "SAVING_ERROR"
	Response   ErrorType = "RESPONSE_ERROR"
	Database   ErrorType = "DATABASE_ERROR"
	Business   ErrorType = "BUSINESS_ERROR"
)

func (e ErrorType) String() string {
	return string(e)
}

type CustomError struct {
	Type      ErrorType
	Domain    string
	Message   string
	CallerErr string
	Extras    map[string]string
}

func NewError(errorType ErrorType, domain string, message string, callerErr string) *CustomError {
	return &CustomError{
		Type:      errorType,
		Domain:    domain,
		Message:   message,
		CallerErr: callerErr,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] [%s]. Message: %s. Error: %s", e.Type, e.Domain, e.Message, e.CallerErr)
}
