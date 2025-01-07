package shared

import (
	"github.com/elmasrisaer/GoSDKTest/internal/clients/rest/httptransport"
)

type CelitechResponse[T any] struct {
	Data     T
	Metadata CelitechResponseMetadata
}

type CelitechResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewCelitechResponse[T any](resp *httptransport.Response[T]) *CelitechResponse[T] {
	return &CelitechResponse[T]{
		Data: resp.Data,
		Metadata: CelitechResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

func (r *CelitechResponse[T]) GetData() T {
	return r.Data
}
