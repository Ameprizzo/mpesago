package mpesa

import "context"

type DisburseRequest struct {
}

type disburser interface {
	Disburse(ctx context.Context, mode Mode, request DisburseRequest) (DisburseResponse, error)
}

type DisburseFunc func(ctx context.Context, mode Mode, request DisburseRequest) (DisburseResponse, error)
