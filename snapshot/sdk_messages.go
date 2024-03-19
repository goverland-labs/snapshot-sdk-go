package snapshot

import (
	"context"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/goverland-labs/snapshot-sdk-go/client"
)

const (
	defaultListMessageLimit   = 500
	listMessageOrderBy        = "mci"
	listMessageOrderDirection = client.OrderDirectionAsc
)

var defaultListMessageOptions = ListMessageOptions{
	Limit:  defaultListMessageLimit,
	Offset: 0,
}

type ListMessageOptions struct {
	LastMCI      int
	Limit        int
	Offset       int
	interceptors []clientv2.RequestInterceptor
}

type ListMessageOption func(options *ListMessageOptions)

func ListMessageWithPagination(limit, offset int) ListMessageOption {
	return func(options *ListMessageOptions) {
		options.Limit = limit
		options.Offset = offset
	}
}

func ListMessageWithMCIFilter(mci int) ListMessageOption {
	return func(options *ListMessageOptions) {
		options.LastMCI = mci
	}
}

func ListMessageWithInterceptors(interceptors []clientv2.RequestInterceptor) ListMessageOption {
	return func(options *ListMessageOptions) {
		options.interceptors = interceptors
	}
}

func (s *SDK) ListMessage(ctx context.Context, opts ...ListMessageOption) ([]*client.MessageFragment, error) {
	options := defaultListMessageOptions
	for _, opt := range opts {
		opt(&options)
	}

	list, err := wrapError(s.client.ListMessages(ctx, int64(options.LastMCI), int64(options.Offset), int64(options.Limit), listMessageOrderBy, listMessageOrderDirection, options.interceptors...))
	if err != nil {
		return nil, err
	}

	return list.GetMessages(), nil
}
