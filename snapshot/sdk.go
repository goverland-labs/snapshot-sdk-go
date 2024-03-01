package snapshot

import (
	"net/http"

	"github.com/goverland-labs/snapshot-sdk-go/client"
)

const (
	ProductionHub      = "https://hub.snapshot.org/graphql"
	ProductionScoreURL = "https://score.snapshot.org/"
	ProductionSeqURL   = "https://seq.snapshot.org/"
)

type SDK struct {
	client client.SnapshotClient

	// TODO setup options
	scoreURL   string
	seqURL     string
	httpClient *http.Client
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
		client:     client.NewClient(options.httpClient, options.baseURL, options.options, interceptors...),
		scoreURL:   ProductionScoreURL,
		seqURL:     ProductionSeqURL,
		httpClient: options.httpClient,
	}
}
