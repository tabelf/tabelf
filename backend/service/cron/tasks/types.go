package tasks

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type JobContext struct {
	Context context.Context
	Client  *resty.Client
}

type JobContextOption func(options *JobContext)

func WithRestyClient(s *resty.Client) JobContextOption {
	return func(options *JobContext) {
		options.Client = s
	}
}

func NewJobContext(ctx context.Context, opts ...JobContextOption) JobContext {
	jobContext := JobContext{Context: ctx, Client: nil}
	for _, o := range opts {
		o(&jobContext)
	}
	return jobContext
}
