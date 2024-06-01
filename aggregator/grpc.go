package main

import "github.com/lucaliebenberg/tolling/types"

type GRPCServer struct {
	types.UnimplementedAggregatorServer
}

func (s *GRPCServer) AggregateDistance(req *types.AggregateRequest) error {
	return nil
}
