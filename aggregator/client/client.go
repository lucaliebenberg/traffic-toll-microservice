package client

import (
	"context"

	"github.com/lucaliebenberg/tolling/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
