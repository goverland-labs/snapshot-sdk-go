package snapshot

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	IDs            []*string
	Limit          int
	Offset         int
	OrderBy        string
	OrderDirection client.OrderDirection
	interceptors   []clientv2.RequestInterceptor
	CreatedAfter   int64
}

type ListProposalOption func(*ListProposalOptions)

type ListProposalPagination struct {
	Limit  int
	Offset int
}

func ListProposalWithPagination(limit, offset int) ListProposalOption {
	return func(options *ListProposalOptions) {
		options.Limit = limit
		options.Offset = offset
	}
}

type ListProposalOrderBy struct {
	OrderBy        string
	OrderDirection client.OrderDirection
}

func ListProposalWithOrderBy(orderBy string, orderDirection client.OrderDirection) ListProposalOption {
	return func(options *ListProposalOptions) {
		options.OrderBy = orderBy
		options.OrderDirection = orderDirection
	}
}

func ListProposalWithSpacesFilter(spaceID ...string) ListProposalOption {
	return func(options *ListProposalOptions) {
		ids := make([]*string, 0, len(spaceID))
		for _, id := range spaceID {
			ids = append(ids, helpers.Ptr(id))
		}

		options.SpaceIDs = ids
	}
}

func ListProposalWithIDFilter(proposalID ...string) ListProposalOption {
	return func(options *ListProposalOptions) {
		ids := make([]*string, 0, len(proposalID))
		for _, id := range proposalID {
			ids = append(ids, helpers.Ptr(id))
		}

		options.IDs = ids
	}
}

func ListProposalCreatedAfter(t time.Time) ListProposalOption {
	return func(options *ListProposalOptions) {
		options.CreatedAfter = 0
		if !t.IsZero() {
			options.CreatedAfter = t.Unix()
		}
	}
}

func ListProposalWithInterceptors(interceptors []clientv2.RequestInterceptor) ListProposalOption {
	return func(options *ListProposalOptions) {
		options.interceptors = interceptors
	}
}

func (s *SDK) ListProposal(ctx context.Context, opts ...ListProposalOption) ([]*client.ProposalFragment, error) {
	options := defaultListProposalOptions
	for _, opt := range opts {
		opt(&options)
	}

	list, err := wrapError(s.client.ListProposals(
		ctx,
		int64(options.Offset),
		int64(options.Limit),
		options.CreatedAfter,
		options.SpaceIDs,
		options.IDs,
		options.OrderBy,
		options.OrderDirection,
		options.interceptors...,
	))
	if err != nil {
		return list.GetProposals(), err
	}

	return list.GetProposals(), nil
}

func (s *SDK) GetProposalByID(ctx context.Context, id string, opts ...clientv2.RequestInterceptor) (*client.ProposalFragment, error) {
	resp, err := wrapError(s.client.ProposalByID(ctx, id, opts...))
	if err != nil {
		return nil, err
	}

	if len(resp.GetProposals()) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrProposalNotFound, id)
	}

	return resp.GetProposals()[0], nil
}
