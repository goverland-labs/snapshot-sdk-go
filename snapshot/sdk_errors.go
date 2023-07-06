package snapshot

import (
	"errors"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

var ErrTooManyRequests = errors.New("too many requests")

func wrapError[T any](resp T, err error) (T, error) {
	if err == nil {
		return resp, err
	}

	e, ok := err.(*clientv2.ErrorResponse)
	if !ok {
		return resp, err
	}

	if !e.HasErrors() {
		return resp, err
	}

	if e.NetworkError == nil {
		return resp, err
	}

	switch e.NetworkError.Code {
	case http.StatusTooManyRequests:
		return resp, ErrTooManyRequests
	}

	return resp, err
}
