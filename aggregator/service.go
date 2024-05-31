package main

import (
	"fmt"

	"github.com/lucaliebenberg/tolling/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
	DistanceSum(int) (float64, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
	// DistanceSum(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	// logrus.WithFields(logrus.Fields{
	// 	"obuid":    distance.OBUID,
	// 	"distance": distance.Value,
	// 	"unix":     distance.Unix,
	// }).Info("aggregating distance")
	fmt.Println("processing and inserting distance in the storage: ", distance)
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) DistanceSum(obuID int) (float64, error) {
	// fmt.Println("processing and inserting distance in the storage: ", distance)
	return i.store.Get(obuID)
}
