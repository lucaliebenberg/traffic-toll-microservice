package main

import (
	"fmt"

	"github.com/lucaliebenberg/tolling/types"
)

type Aggregator interface {
	AggregateDistnace(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistnace(distance types.Distance) error {
	fmt.Println("processing and inserting distance into storage; ", distance)
	return i.store.Insert(distance)
}
