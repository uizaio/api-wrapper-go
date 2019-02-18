package uiza

import "encoding/json"

// ErrorMessage is the message that will show to user.
type ErrorMessage string

type ErrorCode int

// List of values that ErrorType can take.
const (
	ErrorMessageBadRequest         ErrorMessage = "The request was unacceptable, often due to missing a required parameter."
	ErrorMessageAuthentication     ErrorMessage = "No valid API key provided."
	ErrorMessageNotFound           ErrorMessage = "The requested resource doesn't exist."
	ErrorMessageUnProcessable      ErrorMessage = "The syntax of the request is correct (often cause of wrong parameter)."
	ErrorMessageInternalServer     ErrorMessage = "We had a problem with our server. Try again later."
	ErrorMessageServiceUnavailable ErrorMessage = "The server is overloaded or down for maintenance."
	ErrorMessageClient             ErrorMessage = "The error seems to have been caused by the client."
	ErrorMessageServer             ErrorMessage = "The server is aware that it has encountered an error."
)

// List of values that ErrorCode can take.
const (
	ErrorCodeBadRequest         ErrorCode = 400
	ErrorCodeAuthentication     ErrorCode = 401
	ErrorCodeNotFound           ErrorCode = 404
	ErrorCodeUnProcessable      ErrorCode = 422
	ErrorCodeInternalServer     ErrorCode = 500
	ErrorCodeServiceUnavailable ErrorCode = 503
)

// Error is the response returned when a call is unsuccessful.
// For more details see  https://docs.uiza.io/#errors-code.
type Error struct {
	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong.
	Err            error        `json:"-"`
	HTTPStatusCode int          `json:"-"`
	Code           ErrorCode    `json:"code,omitempty"`
	Message        ErrorMessage `json:"-"`
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
	e.uizaErr.Message = ErrorMessageAuthentication
	return e.uizaErr.Error()
}

type BadRequestError struct {
	uizaErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *BadRequestError) Error() string {
	e.uizaErr.DescriptionLink = "https://docs.uiza.io/#errors-code"
	e.uizaErr.Message = ErrorMessageBadRequest
	return e.uizaErr.Error()
}

type UnProcessableError struct {
	uizaErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *UnProcessableError) Error() string {
	e.uizaErr.DescriptionLink = "https://docs.uiza.io/#errors-code"
	e.uizaErr.Message = ErrorMessageUnProcessable
	return e.uizaErr.Error()
}

type NotFoundError struct {
	uizaErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *NotFoundError) Error() string {
	e.uizaErr.DescriptionLink = "https://docs.uiza.io/#errors-code"
	e.uizaErr.Message = ErrorMessageNotFound
	return e.uizaErr.Error()
}

type InternalServerError struct {
	uizaErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *InternalServerError) Error() string {
	e.uizaErr.DescriptionLink = "https://docs.uiza.io/#errors-code"
	e.uizaErr.Message = ErrorMessageInternalServer
	return e.uizaErr.Error()
}

type ServiceUnavailableError struct {
	uizaErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *ServiceUnavailableError) Error() string {
	e.uizaErr.DescriptionLink = "https://docs.uiza.io/#errors-code"
	e.uizaErr.Message = ErrorMessageServiceUnavailable
	return e.uizaErr.Error()
}
