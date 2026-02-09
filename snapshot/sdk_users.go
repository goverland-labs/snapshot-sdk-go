package snapshot

import (
	"context"
	"errors"

	"github.com/goverland-labs/snapshot-sdk-go/client"
)

var (
	ErrNoAddressesProvided = errors.New("no addresses provided")
)

func (s *SDK) ListUsers(
	ctx context.Context,
	addresses []string,
) ([]*client.UserFragment, error) {
	if len(addresses) == 0 {
		return nil, ErrNoAddressesProvided
	}

	list, err := wrapError(s.client.ListUsers(
		ctx,
		addresses,
	))
	if err != nil {
		return nil, err
	}

	return list.GetUsers(), nil
}
