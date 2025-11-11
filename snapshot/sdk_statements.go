package snapshot

import (
	"context"
	"errors"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/goverland-labs/snapshot-sdk-go/client"
)

var (
	ErrNoDelegatesProvided = errors.New("no delegates provided")
)

const defaultListStatementsLimit = 1000

var defaultListStatementsOptions = ListStatementsOptions{
	Limit: defaultListStatementsLimit,
}

type ListStatementsOptions struct {
	Limit        int
	Offset       int
	interceptors []clientv2.RequestInterceptor
}

type ListStatementsOption func(options *ListStatementsOptions)

func ListStatementsWithPagination(limit, offset int) ListStatementsOption {
	return func(options *ListStatementsOptions) {
		options.Limit = limit
		options.Offset = offset
	}
}

func (s *SDK) ListStatements(
	ctx context.Context,
	spaceID string,
	delegatesAddresses []string,
	opts ...ListStatementsOption,
) ([]*client.StatementFragment, error) {
	options := defaultListStatementsOptions
	for _, opt := range opts {
		opt(&options)
	}

	if len(delegatesAddresses) == 0 {
		return nil, ErrNoDelegatesProvided
	}

	list, err := wrapError(s.client.ListStatements(
		ctx,
		spaceID,
		delegatesAddresses,
		int64(options.Offset),
		int64(options.Limit),
		options.interceptors...,
	))
	if err != nil {
		return nil, err
	}

	return list.GetStatements(), nil
}
