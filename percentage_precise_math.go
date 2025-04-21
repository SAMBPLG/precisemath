package precisemath

import (
	"github.com/shopspring/decimal"
)

// PercentageFromStringDecimal returns decimal.Decimal s from v
func PercentageFromStringDecimal(s, v string) decimal.Decimal {
	subamount := StringToFloat64Decimal(s)
	amount := StringToFloat64Decimal(v)
	if amount.IsZero() {
		return decimal.Zero
	}
	return subamount.Div(amount).Mul(hundred)
}

// PercentageFromString returns high-precision percentage of s from v
func PercentageFromString(s, v string) float32 {
	p := PercentageFromStringDecimal(s, v)
	if p.IsZero() {
		return 0
	}
	return Float64ToFloat32(p.InexactFloat64())
}

// PercentageFromStringRound returns high-precision percentage of s from v rounded upto round
func PercentageFromStringRound(of, from string, round int32) float32 {
	p := PercentageFromStringDecimal(of, from)
	if p.IsZero() {
		return 0
	}
	return Float64ToFloat32(p.Round(round).InexactFloat64())
}

// Percentage returns high-precision percentage of s from v
func Percentage(s, v float64) float32 {
	a := Float64ToString(s)
	b := Float64ToString(v)
	return PercentageFromString(a, b)
}

// PercentageRound returns high-precision percentage of s from v
func PercentageRound(s, v float64, decimalLength int32) float32 {
	a := Float64ToString(s)
	b := Float64ToString(v)
	return PercentageFromStringRound(a, b, decimalLength)
}

// AmountOfDecimal returns decimal.Decimal of s percentage from v
func AmountOfDecimal(s float64, v float64) decimal.Decimal {
	a := Float64ToString(s)
	b := Float64ToString(v)
	percent, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	from, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	return percent.Div(hundred).Mul(from)
}

// AmountOf returns s percentage value from v
func AmountOf(s float64, v float64) float64 {
	p := AmountOfDecimal(s, v)
	if p.IsZero() {
		return 0
	}
	return p.InexactFloat64()
}

// AmountOfRound returns s percentage value from v and rounded upto round
func AmountOfRound(s float64, v float64, round int32) float64 {
	p := AmountOfDecimal(s, v)
	if p.IsZero() {
		return 0
	}
	return p.Round(round).InexactFloat64()
}
