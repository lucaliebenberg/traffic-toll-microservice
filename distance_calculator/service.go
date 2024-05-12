package main

import (
	"fmt"

	"github.com/lucaliebenberg/tolling/types"
)

// We like to end our interfaces with (er)
type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculateService struct{}

func NewCalculatorService() CalculatorServicer {
	return &CalculateService{}
}

func (s *CalculateService) CalculateDistance(data types.OBUData) (float64, error) {
	fmt.Println("calculating the distance")
	return 0.0, nil
}
