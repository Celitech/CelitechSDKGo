package shared

import (
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
)

type CelitechError struct {
	Err      error
	Body     []byte
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
		Metadata: CelitechErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *CelitechError) Error() string {
	return e.Err.Error()
}
