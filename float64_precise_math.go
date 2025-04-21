package precisemath

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// Float64 casts float64 into high-precision float64
func Float64(of float64) float64 {
	res, err := decimal.NewFromString(Float64ToString(of))
	if err != nil {
		panic(err)
	}
	return res.InexactFloat64()
}

// Float64ToString converts float64 to string representation
func Float64ToString(of float64) string {
	return strconv.FormatFloat(of, 'g', -1, 64)
}

// Float64ToStringScientific converts float64 into scientific notation string
func Float64ToStringScientific(of float64) string {
	return strconv.FormatFloat(of, 'e', -1, 64)
}

// Float64Round returns float64 rounded upto the round. Expect this function to operate with high-precision calculation
func Float64Round(of float64, round int32) float64 {
	res, err := decimal.NewFromString(Float64ToString(of))
	if err != nil {
		panic(err)
	}
	return res.Round(round).InexactFloat64()
}

// Float64ToFloat32 converts float64 into float32. Expect some precision loss since float32 only has ~7 significant digits of precision
func Float64ToFloat32(of float64) float32 {
	return StringToFloat32(Float64ToString(of))
}

// Float64DivDecimal returns decimal.Decimal of s.Div(v)
func Float64DivDecimal(s, v float64) decimal.Decimal {
	aa, err := decimal.NewFromString(Float64ToString(s))
	if err != nil {
		panic(err)
	}
	bb, err := decimal.NewFromString(Float64ToString(v))
	if err != nil {
		panic(err)
	}
	if bb.IsZero() {
		return decimal.Zero
	}
	return aa.Div(bb)
}

// Float64SubDecimal returns decimal.Decimal of s - v
func Float64SubDecimal(s, v float64) decimal.Decimal {
	aa, err := decimal.NewFromString(Float64ToString(s))
	if err != nil {
		panic(err)
	}
	bb, err := decimal.NewFromString(Float64ToString(v))
	if err != nil {
		panic(err)
	}
	return aa.Sub(bb)
}

// Float64Sub returns high-precision subtitution operation of s - v
func Float64Sub(s, v float64) float64 {
	return Float64SubDecimal(s, v).InexactFloat64()
}

// Float64SubRound returns high-precision subtitution operation of s - v rounded up to round
func Float64SubRound(s, v float64, round int32) float64 {
	return Float64SubDecimal(s, v).Round(round).InexactFloat64()
}

// Float64SumDecimal returns decimal.Decimal of summary operation from s
func Float64SumDecimal(s ...float64) decimal.Decimal {
	v := decimal.Zero
	for _, f := range s {
		v = v.Add(StringToFloat64Decimal(Float64ToString(f)))
	}
	return v
}

// Float64Sum returns high-precision float64 of summary operation from s
func Float64Sum(s ...float64) float64 {
	p := Float64SumDecimal(s...)
	return p.InexactFloat64()
}

// Float64SumRound returns high-precision float64 of summary operation from s rounded upto round
func Float64SumRound(round int32, s ...float64) float64 {
	p := Float64SumDecimal(s...)
	return p.Round(round).InexactFloat64()
}
