package snapshot

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Options struct {
	baseURL      string
	httpClient   *http.Client
	apiKey       string
	interceptors []clientv2.RequestInterceptor
	options      *clientv2.Options
}

type Option func(opts *Options)

func WithBaseURL(baseURL string) Option {
	return func(opts *Options) {
		opts.baseURL = baseURL
	}
}

func WithApiKey(apiKey string) Option {
	return func(opts *Options) {
		opts.apiKey = apiKey
	}
}

func WithInterceptors(interceptors []clientv2.RequestInterceptor) Option {
	return func(opts *Options) {
		opts.interceptors = interceptors
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(opts *Options) {
		opts.httpClient = client
	}
}

func WithOptions(options *clientv2.Options) Option {
	return func(opts *Options) {
		opts.options = options
	}
}

func ApiKeyInterceptor(apiKey string) clientv2.RequestInterceptor {
	return func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
		req.Header.Set("x-api-key", apiKey)

		return next(ctx, req, gqlInfo, res)
	}
}
