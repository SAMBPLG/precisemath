package precisemath

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func StringToFloat64(of string) float64 {
	a, err := strconv.ParseFloat(of, 64)
	if err != nil {
		panic(err)
	}
	return Float64(a)
}

func StringToFloat64Decimal(of string) decimal.Decimal {
	a, err := strconv.ParseFloat(of, 64)
	if err != nil {
		panic(err)
	}
	return decimal.NewFromFloat(a)
}

func StringToFloat64Round(of string, round int32) float64 {
	a := StringToFloat64(of)
	return Float64Round(a, round)
}

func StringToFloat32(of string) float32 {
	a, err := strconv.ParseFloat(of, 32)
	if err != nil {
		panic(err)
	}
	return Float32(float32(a))
}

func StringToFloat32Round(of string, round int32) float32 {
	a, err := strconv.ParseFloat(of, 32)
	if err != nil {
		panic(err)
	}
	return Float32Round(float32(a), round)
}
