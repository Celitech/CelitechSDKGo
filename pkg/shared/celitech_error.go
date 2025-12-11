package shared

import (
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"net/http"
)

type CelitechError struct {
	Err      error
	Body     []byte
	Raw      *http.Response
	Metadata CelitechErrorMetadata
}

type CelitechErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewCelitechError[T any](transportError *httptransport.ErrorResponse[T]) *CelitechError {
	return &CelitechError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: CelitechErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *CelitechError) Error() string {
	return e.Err.Error()
}
