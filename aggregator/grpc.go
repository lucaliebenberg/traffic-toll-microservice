package main

import (
	"context"

	"github.com/lucaliebenberg/tolling/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

// TRANSPORT LAYER
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> types.Distance
// Webpack -> types.Webpack -> types>Distance

// BUSINESS LAYER
// business layer type (main type everyone needs to conver to)

func (s *GRPCAggregatorServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
