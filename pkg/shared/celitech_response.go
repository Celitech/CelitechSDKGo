package shared

import (
	"encoding/json"
	"github.com/Celitech/CelitechSDKGo/internal/clients/rest/httptransport"
	"net/http"
)

type CelitechResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata CelitechResponseMetadata
}

type CelitechResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

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

func (r *CelitechResponse[T]) GetData() T {
	return r.Data
}

func (r CelitechResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: CelitechResponse to string"
	}
	return string(jsonData)
}
