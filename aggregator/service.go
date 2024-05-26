package main

import (
	"github.com/lucaliebenberg/tolling/types"
	"github.com/sirupsen/logrus"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
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
	logrus.WithFields(logrus.Fields{
		"obuid":    distance.OBUID,
		"distance": distance.Value,
		"unix":     distance.Unix,
	}).Info("aggregating distance")
	return i.store.Insert(distance)
}
