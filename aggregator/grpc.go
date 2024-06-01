package main

import "github.com/lucaliebenberg/tolling/types"

type GRPCServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewGRPCServer(svc Aggregator) *GRPCServer {
	return &GRPCServer{
		svc: svc,
	}
}

// TRANSPORT LAYER
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> types.Distance
// Webpack -> types.Webpack -> types>Distance

// BUSINESS LAYER
// business layer type (main type everyone needs to conver to)

func (s *GRPCServer) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return s.svc.AggregateDistance(distance)
}
