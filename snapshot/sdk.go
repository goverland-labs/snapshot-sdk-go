package snapshot

import (
	"net/http"

	"github.com/goverland-labs/sdk-snapshot-go/client"
)

const ProductionHub = "https://hub.snapshot.org/graphql"

type SDK struct {
	client client.SnapshotClient
}

func NewSDK(opts ...Option) *SDK {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.baseURL == "" {
		options.baseURL = ProductionHub
	}

	if options.httpClient == nil {
		options.httpClient = http.DefaultClient
	}

	interceptors := options.interceptors
	if options.apiKey != "" {
		interceptors = append(interceptors, ApiKeyInterceptor(options.apiKey))
	}

	return &SDK{
		client: client.NewClient(options.httpClient, options.baseURL, options.options, interceptors...),
	}
}
