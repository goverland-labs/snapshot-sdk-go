package snapshot

import (
	"context"
	"errors"
	"fmt"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/goverland-labs/sdk-snapshot-go/client"
	"github.com/goverland-labs/sdk-snapshot-go/helpers"
)

var ErrSpaceNotFound = errors.New("space not found")

const defaultListSpaceLimit = 500

var defaultListSpaceOptions = ListSpaceOptions{
	Limit:  defaultListSpaceLimit,
	Offset: 0,
}

type ListSpaceOptions struct {
	IDs          []*string
	Limit        int
	Offset       int
	interceptors []clientv2.RequestInterceptor
}

type ListSpaceOption interface {
	Apply(options *ListSpaceOptions)
}

type ListSpacePagination struct {
	Limit  int
	Offset int
}

func ListSpaceWithPagination(limit, offset int) ListSpacePagination {
	return ListSpacePagination{
		Limit:  limit,
		Offset: offset,
	}
}

func (o ListSpacePagination) Apply(options *ListSpaceOptions) {
	options.Limit = o.Limit
	options.Offset = o.Offset
}

type ListSpaceIDs struct {
	IDs []string
}

func ListSpaceWithIDs(id ...string) ListSpaceIDs {
	return ListSpaceIDs{
		IDs: id,
	}
}

func (o ListSpaceIDs) Apply(options *ListSpaceOptions) {
	ids := make([]*string, 0, len(o.IDs))
	for _, id := range o.IDs {
		ids = append(ids, helpers.Ptr(id))
	}

	options.IDs = ids
}

type ListSpaceInterceptors struct {
	Interceptors []clientv2.RequestInterceptor
}

func (o ListSpaceInterceptors) Apply(options *ListSpaceOptions) {
	options.interceptors = o.Interceptors
}

func (s *SDK) ListSpace(ctx context.Context, opts ...ListSpaceOption) ([]*client.SpaceFragment, error) {
	options := defaultListSpaceOptions
	for _, opt := range opts {
		opt.Apply(&options)
	}

	list, err := s.client.ListSpaces(ctx, int64(options.Offset), int64(options.Limit), options.IDs, options.interceptors...)
	if err != nil {
		return nil, err
	}

	return list.GetSpaces(), nil
}

func (s *SDK) GetSpaceByID(ctx context.Context, id string, opts ...clientv2.RequestInterceptor) (*client.SpaceFragment, error) {
	resp, err := s.client.SpaceByID(ctx, id, opts...)
	if err != nil {
		return nil, err
	}

	if len(resp.GetSpaces()) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrSpaceNotFound, id)
	}

	return resp.GetSpaces()[0], nil
}
