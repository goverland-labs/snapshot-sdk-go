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

type Option interface {
	Apply(opts *Options)
}

type WithBaseURL struct {
	BaseURL string
}

func (o WithBaseURL) Apply(options *Options) {
	options.baseURL = o.BaseURL
}

type WithInterceptors struct {
	Interceptors []clientv2.RequestInterceptor
}

func (o WithInterceptors) Apply(options *Options) {
	options.interceptors = o.Interceptors
}

type WithHTTPClient struct {
	Client *http.Client
}

func (o WithHTTPClient) Apply(options *Options) {
	options.httpClient = o.Client
}

type WithOptions struct {
	Options *clientv2.Options
}

func (o WithOptions) Apply(options *Options) {
	options.options = o.Options
}
