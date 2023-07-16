package snapshot

import (
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Options struct {
	baseURL      string
	httpClient   *http.Client
	interceptors []clientv2.RequestInterceptor
	options      *clientv2.Options
}

type Option func(opts *Options)

func WithBaseURL(baseURL string) Option {
	return func(opts *Options) {
		opts.baseURL = baseURL
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
