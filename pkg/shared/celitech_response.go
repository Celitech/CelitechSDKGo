package shared

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"net/http"
)

// CelitechResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type CelitechResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata CelitechResponseMetadata
}

// CelitechResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type CelitechResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewCelitechResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewCelitechResponse[T any](resp *httptransport.Response[T]) *CelitechResponse[T] {
	return &CelitechResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: CelitechResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *CelitechResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r CelitechResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: CelitechResponse to string"
	}
	return string(jsonData)
}
