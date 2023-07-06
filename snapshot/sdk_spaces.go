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
const defaultRankingLimit = 20

var defaultListSpaceOptions = ListSpaceOptions{
	Limit:  defaultListSpaceLimit,
	Offset: 0,
}

var defaultRankingOptions = RankingOptions{
	Limit:  defaultRankingLimit,
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

type RankingOptions struct {
	Category     string
	Network      string
	Limit        int
	Offset       int
	interceptors []clientv2.RequestInterceptor
}

type RankingOption interface {
	Apply(options *RankingOptions)
}

type RankingPagination struct {
	Limit  int
	Offset int
}

func RankingWithPagination(limit, offset int) RankingPagination {
	return RankingPagination{
		Limit:  limit,
		Offset: offset,
	}
}

func (o RankingPagination) Apply(options *RankingOptions) {
	options.Limit = o.Limit
	options.Offset = o.Offset
}

type RankingCategoryOption struct {
	Category string
}

func RankingWithCategory(cagegory string) RankingCategoryOption {
	return RankingCategoryOption{
		Category: cagegory,
	}
}

func (o RankingCategoryOption) Apply(options *RankingOptions) {
	options.Category = o.Category
}

type RankingNetworkOption struct {
	Network string
}

func RankingWithNetwork(network string) RankingNetworkOption {
	return RankingNetworkOption{
		Network: network,
	}
}

func (o RankingNetworkOption) Apply(options *RankingOptions) {
	options.Network = o.Network
}

type RankingInterceptors struct {
	Interceptors []clientv2.RequestInterceptor
}

func (o RankingInterceptors) Apply(options *RankingOptions) {
	options.interceptors = o.Interceptors
}

func (s *SDK) ListSpace(ctx context.Context, opts ...ListSpaceOption) ([]*client.SpaceFragment, error) {
	options := defaultListSpaceOptions
	for _, opt := range opts {
		opt.Apply(&options)
	}

	list, err := wrapError(s.client.ListSpaces(ctx, int64(options.Offset), int64(options.Limit), options.IDs, options.interceptors...))
	if err != nil {
		return nil, err
	}

	return list.GetSpaces(), nil
}

func (s *SDK) GetRanking(ctx context.Context, opts ...RankingOption) ([]string, error) {
	options := defaultRankingOptions
	for _, opt := range opts {
		opt.Apply(&options)
	}

	list, err := wrapError(s.client.ListRanking(ctx, int64(options.Offset), int64(options.Limit), options.Category, options.Network, options.interceptors...))
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(list.GetRanking().GetItems()))
	for _, item := range list.GetRanking().GetItems() {
		id := item.GetID()
		if id == "" {
			continue
		}

		ids = append(ids, item.GetID())
	}

	return ids, nil
}

func (s *SDK) GetSpaceByID(ctx context.Context, id string, opts ...clientv2.RequestInterceptor) (*client.SpaceFragment, error) {
	resp, err := wrapError(s.client.SpaceByID(ctx, id, opts...))
	if err != nil {
		return nil, err
	}

	if len(resp.GetSpaces()) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrSpaceNotFound, id)
	}

	return resp.GetSpaces()[0], nil
}
