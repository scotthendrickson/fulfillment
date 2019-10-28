package epFulfillment

import "fmt"

//While working on this project EasyPost released a GoLang client for their shipping API (https://github.com/EasyPost/easypost-go). The error handling on that project was
//very useful and so I borrowed it. You can reference that work here https://github.com/EasyPost/easypost-go/blob/master/error.go.

// APIError provides data on why an API request failed. It will be the type
// returned by Client methods when the HTTP API returns an HTTP error code. It
// will not be returned on network, parsing or context errors.
type APIError struct {
	// Status is the HTTP text status of the response.
	Status string
	// StatusCode is the HTTP numerical status code of the response.
	StatusCode int
	// Code is an error code returned by the API. It may be empty.
	Code string `json:"code,omitempty"`
	// Message is a human-readable error code returned by the API. It may be
	// empty.
	Message string `json:"message,omitempty"`
	// Field may be provided when the error relates to a specific field.
	Field string `json:"field,omitempty"`
	// Suggestion may be provided if the API can provide a suggestion to fix
	// the error.
	Suggestion string `json:"suggestion,omitempty"`
	// Errors may be provided if there are multiple errors, for example if
	// multiple fields have invalid values.
	Errors []*APIError `json:"errors,omitempty"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		if e.Code != "" {
			return e.Code + " " + e.Message
		}
		return e.Message
	}
	if e.Code != "" {
		return e.Code
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.Status)
}

type apiErrorResponse struct {
	Error *APIError `json:"errors,omitempty"`
}
