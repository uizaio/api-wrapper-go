package uiza

import "encoding/json"

// ErrorType is the list of allowed values for the error's type.
type ErrorType string

// List of values that ErrorType can take.
const (
	ErrorTypeAuthentication ErrorType = "Unauthorized"
)

// Error is the response returned when a call is unsuccessful.
// For more details see  https://docs.uiza.io/#errors-code.
type Error struct {
	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong. For example, Err may hold a CardError that indicates
	// exactly what went wrong during charging a card.
	Err error `json:"-"`
	// HTTPStatusCode  int       `json:"status,omitempty"`
	Code int       `json:"code,omitempty"`
	Type ErrorType `json:"type"`
	// Data            string    `json:"data,omitempty"`
	// Retryable       bool      `json:"retryable,omitempty"`
	// Message         string    `json:"message,omitempty"`
	// Version         int       `json:"version"`
	// DateTime        string    `json:"datetime,omitempty"`
	// Policy          string    `json:"policy,omitempty"`
	// RequestID       string    `json:"requestId,omitempty"`
	// ServiceName     string    `json:"serviceName,omitempty"`
	DescriptionLink string
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// AuthenticationError is a failure to properly authenticate during a request.
type AuthenticationError struct {
	uizaErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *AuthenticationError) Error() string {
	e.uizaErr.DescriptionLink = "https://docs.uiza.io/#authentication"
	return e.uizaErr.Error()
}
