package snapshot

import (
	"context"
	"errors"
	"fmt"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/goverland-labs/sdk-snapshot-go/client"
	"github.com/goverland-labs/sdk-snapshot-go/helpers"
)

var ErrProposalNotFound = errors.New("proposal not found")

const (
	defaultListProposalLimit          = 500
	defaultListProposalOrderBy        = "created"
	defaultListProposalOrderDirection = client.OrderDirectionDesc
)

var defaultListProposalOptions = ListProposalOptions{
	Limit:          defaultListProposalLimit,
	Offset:         0,
	OrderBy:        defaultListProposalOrderBy,
	OrderDirection: defaultListProposalOrderDirection,
}

type ListProposalOptions struct {
	SpaceIDs       []*string
	Limit          int
	Offset         int
	OrderBy        string
	OrderDirection client.OrderDirection
	interceptors   []clientv2.RequestInterceptor
}

type ListProposalOption interface {
	Apply(options *ListProposalOptions)
}

type ListProposalPagination struct {
	Limit  int
	Offset int
}

func ListProposalWithPagination(limit, offset int) ListProposalPagination {
	return ListProposalPagination{
		Limit:  limit,
		Offset: offset,
	}
}

func (o ListProposalPagination) Apply(options *ListProposalOptions) {
	options.Limit = o.Limit
	options.Offset = o.Offset
}

type ListProposalOrderBy struct {
	OrderBy        string
	OrderDirection client.OrderDirection
}

func ListProposalWithOrderBy(orderBy string, orderDirection client.OrderDirection) ListProposalOrderBy {
	return ListProposalOrderBy{
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

func (o ListProposalOrderBy) Apply(options *ListProposalOptions) {
	options.OrderBy = o.OrderBy
	options.OrderDirection = o.OrderDirection
}

type ListProposalWithSpacesFilter struct {
	SpaceIDs []string
}

func ListProposalWithSpaces(id ...string) ListProposalWithSpacesFilter {
	return ListProposalWithSpacesFilter{
		SpaceIDs: id,
	}
}

func (o ListProposalWithSpacesFilter) Apply(options *ListProposalOptions) {
	ids := make([]*string, 0, len(o.SpaceIDs))
	for _, id := range o.SpaceIDs {
		ids = append(ids, helpers.Ptr(id))
	}

	options.SpaceIDs = ids
}

type ListProposalInterceptors struct {
	Interceptors []clientv2.RequestInterceptor
}

func (o ListProposalInterceptors) Apply(options *ListProposalOptions) {
	options.interceptors = o.Interceptors
}

func (s *SDK) ListProposal(ctx context.Context, opts ...ListProposalOption) ([]*client.ProposalFragment, error) {
	options := defaultListProposalOptions
	for _, opt := range opts {
		opt.Apply(&options)
	}

	list, err := s.client.ListProposals(
		ctx,
		int64(options.Offset),
		int64(options.Limit),
		options.SpaceIDs,
		options.OrderBy,
		options.OrderDirection,
		options.interceptors...,
	)
	if err != nil {
		return nil, err
	}

	return list.GetProposals(), nil
}

func (s *SDK) GetProposalByID(ctx context.Context, id string, opts ...clientv2.RequestInterceptor) (*client.ProposalFragment, error) {
	resp, err := s.client.ProposalByID(ctx, id, opts...)
	if err != nil {
		return nil, err
	}

	if len(resp.GetProposals()) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrProposalNotFound, id)
	}

	return resp.GetProposals()[0], nil
}
