package snapshot

import (
	"context"

	"github.com/Yamashou/gqlgenc/clientv2"
)

func (s *SDK) ListNetworks(ctx context.Context, opts ...clientv2.RequestInterceptor) ([]string, error) {
	resp, err := wrapError(s.client.ListNetworks(ctx, opts...))
	if err != nil {
		return nil, err
	}

	networks := make([]string, 0, len(resp.GetNetworks()))
	for _, n := range resp.GetNetworks() {
		networks = append(networks, n.GetID())
	}

	return networks, nil
}
