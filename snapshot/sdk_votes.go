package snapshot

import (
	"context"
	"errors"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/goverland-labs/snapshot-sdk-go/client"
	"github.com/goverland-labs/snapshot-sdk-go/helpers"
)

const defaultListVotesLimit = 1000

var defaultListVotesOptions = ListVotesOptions{
	Limit:          defaultListVotesLimit,
	OrderBy:        "created",
	OrderDirection: client.OrderDirectionDesc,
}

type ListVotesOptions struct {
	ProposalIDs    []*string
	Limit          int
	Offset         int
	OrderBy        string
	OrderDirection client.OrderDirection
	interceptors   []clientv2.RequestInterceptor
	CreatedAfter   int64
}

type ListVotesOption func(options *ListVotesOptions)

func ListVotesWithPagination(limit, offset int) ListVotesOption {
	return func(options *ListVotesOptions) {
		options.Limit = limit
		options.Offset = offset
	}
}

func ListVotesWithOrderBy(orderBy string, orderDirection client.OrderDirection) ListVotesOption {
	return func(options *ListVotesOptions) {
		options.OrderBy = orderBy
		options.OrderDirection = orderDirection
	}
}

func ListVotesWithProposalIDsFilter(proposalID ...string) ListVotesOption {
	return func(options *ListVotesOptions) {
		ids := make([]*string, 0, len(proposalID))
		for _, id := range proposalID {
			ids = append(ids, helpers.Ptr(id))
		}

		options.ProposalIDs = ids
	}
}

func ListVotesWithInterceptors(interceptors []clientv2.RequestInterceptor) ListVotesOption {
	return func(options *ListVotesOptions) {
		options.interceptors = interceptors
	}
}

func ListVotesCreatedAfter(t time.Time) ListVotesOption {
	return func(options *ListVotesOptions) {
		options.CreatedAfter = 0
		if !t.IsZero() {
			options.CreatedAfter = t.Unix()
		}
	}
}

func (s *SDK) ListVotes(ctx context.Context, opts ...ListVotesOption) ([]*client.VoteFragment, error) {
	options := defaultListVotesOptions
	for _, opt := range opts {
		opt(&options)
	}

	list, err := wrapError(s.client.ListVotes(ctx, options.ProposalIDs, int64(options.Offset), int64(options.Limit), options.OrderBy, options.OrderDirection, options.CreatedAfter, options.interceptors...))
	if err != nil {
		return nil, err
	}

	return list.GetVotes(), nil
}

func (s *SDK) VoteByID(ctx context.Context, id string) (*client.VoteFragment, error) {
	list, err := wrapError(s.client.VoteByID(ctx, id))
	if err != nil {
		return nil, err
	}

	votes := list.GetVotes()
	if len(votes) == 0 {
		return nil, nil
	}
	if len(votes) > 1 {
		return nil, errors.New("more than one vote returned")
	}

	return votes[0], nil
}
