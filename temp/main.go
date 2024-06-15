package main

import (
	"context"
	"log"
	"time"

	"github.com/lucaliebenberg/tolling/aggregator/client"
	"github.com/lucaliebenberg/tolling/types"
)

func main() {

	c, err := client.NewGRPCClient(":3001")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 55.45,
		Unix:  time.Now().UnixNano(),
	}); err != nil {
		log.Fatal(err)
	}
}
