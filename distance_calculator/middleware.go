package main

import (
	"time"

	"github.com/lucaliebenberg/tolling/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleare(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": dist,
		}).Info("calculate distance")
	}(time.Now())
	dist, err = m.next.CalculateDistance(data)
	return
}
