package snapshot

import (
	"context"
	"errors"
	"fmt"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/goverland-labs/snapshot-sdk-go/client"
	"github.com/goverland-labs/snapshot-sdk-go/helpers"
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

type ListSpaceOption func(options *ListSpaceOptions)

func ListSpaceWithPagination(limit, offset int) ListSpaceOption {
	return func(options *ListSpaceOptions) {
		options.Limit = limit
		options.Offset = offset
	}
}

func ListSpaceWithIDFilter(spaceID ...string) ListSpaceOption {
	return func(options *ListSpaceOptions) {
		ids := make([]*string, 0, len(spaceID))
		for _, id := range spaceID {
			ids = append(ids, helpers.Ptr(id))
		}

		options.IDs = ids
	}
}

func ListSpaceWithInterceptors(interceptors []clientv2.RequestInterceptor) ListSpaceOption {
	return func(options *ListSpaceOptions) {
		options.interceptors = interceptors
	}
}

type RankingOptions struct {
	Category     string
	Network      string
	Limit        int
	Offset       int
	interceptors []clientv2.RequestInterceptor
}

type RankingOption func(options *RankingOptions)

type RankingPagination struct {
	Limit  int
	Offset int
}

func RankingWithPagination(limit, offset int) RankingOption {
	return func(options *RankingOptions) {
		options.Limit = limit
		options.Offset = offset
	}
}

func RankingWithCategory(category string) RankingOption {
	return func(options *RankingOptions) {
		options.Category = category
	}
}

func RankingWithNetwork(network string) RankingOption {
	return func(options *RankingOptions) {
		options.Network = network
	}
}

func RankingWithInterceptors(interceptors []clientv2.RequestInterceptor) RankingOption {
	return func(options *RankingOptions) {
		options.interceptors = interceptors
	}
}

func (s *SDK) ListSpace(ctx context.Context, opts ...ListSpaceOption) ([]*client.SpaceFragment, error) {
	options := defaultListSpaceOptions
	for _, opt := range opts {
		opt(&options)
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
		opt(&options)
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
