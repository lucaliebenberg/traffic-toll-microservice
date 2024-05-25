package main

import (
	"math"

	"github.com/lucaliebenberg/tolling/aggregator/client"
	"github.com/lucaliebenberg/tolling/types"
)

// We like to end our interfaces with (er)
type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculateService struct {
	prevPoint []float64
	aggClient client.Client
}

func NewCalculatorService() CalculatorServicer {
	return &CalculateService{}
}

func (s *CalculateService) CalculateDistance(data types.OBUData) (float64, error) {
	distance := 0.0
	if len(s.prevPoint) > 0 {
		distance = calculateDistance(s.prevPoint[0], s.prevPoint[1], data.Lat, data.Long)
	}
	s.prevPoint = []float64{data.Lat, data.Long}
	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
