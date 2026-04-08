package sharedmodels

import (
	"example.com/celitech/internal/clients/rest/httptransport"
	"net/http"
)

// CelitechError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type CelitechError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata CelitechErrorMetadata
}

// CelitechErrorMetadata contains HTTP metadata associated with an error response.
type CelitechErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewCelitechError creates a new CelitechError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewCelitechError[T any](transportError *httptransport.ErrorResponse[T]) *CelitechError[T] {
	return &CelitechError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: CelitechErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *CelitechError[T]) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As to traverse the chain.
func (e *CelitechError[T]) Unwrap() error {
	return e.Err
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *CelitechError[T]) GetData() *T {
	return e.Data
}
