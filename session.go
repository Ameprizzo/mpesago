package mpesa

import "context"

type session interface {
	Session(ctx context.Context) (response SessionResponse, err error)
}

type SessionFunc func(ctx context.Context) (response SessionResponse, err error)
