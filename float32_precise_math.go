package precisemath

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// Float32 casts float32 into high-precision float32
func Float32(of float32) float32 {
	res, err := decimal.NewFromString(Float32ToString(of))
	if err != nil {
		panic(err)
	}
	return float32(res.InexactFloat64())
}

// Float32ToString converts float32 to string representation
func Float32ToString(of float32) string {
	return strconv.FormatFloat(float64(of), 'g', -1, 32)
}

// Float32ToStringScientific converts float32 into scientific notation string
func Float32ToStringScientific(of float32) string {
	return strconv.FormatFloat(float64(of), 'e', -1, 32)
}

// Float32Round returns float32 rounded upto the round. Expect this function to operate with high-precision calculation
func Float32Round(of float32, round int32) float32 {
	res, err := decimal.NewFromString(Float32ToString(of))
	if err != nil {
		panic(err)
	}
	return float32(res.Round(round).InexactFloat64())
}

// Float32ToFloat64 converts float32 into float64 without precision loss
func Float32ToFloat64(of float32) float64 {
	return StringToFloat64(Float32ToString(of))
}

// Float32DivDecimal returns decimal.Decimal of a.Div(b)
func Float32DivDecimal(s, v float32) decimal.Decimal {
	aa, err := decimal.NewFromString(Float32ToString(s))
	if err != nil {
		panic(err)
	}
	bb, err := decimal.NewFromString(Float32ToString(v))
	if err != nil {
		panic(err)
	}
	if bb.IsZero() {
		return decimal.Zero
	}
	return aa.Div(bb)
}

// Float32SubDecimal returns decimal.Decimal of s - v
func Float32SubDecimal(s, v float32) decimal.Decimal {
	aa, err := decimal.NewFromString(Float32ToString(s))
	if err != nil {
		panic(err)
	}
	bb, err := decimal.NewFromString(Float32ToString(v))
	if err != nil {
		panic(err)
	}
	return aa.Sub(bb)
}

// Float32Sub returns high-precision subtitution operation of s - v
func Float32Sub(s, v float32) float32 {
	res := Float32SubDecimal(s, v).String()
	return StringToFloat32(res)
}

// Float32Sub returns high-precision subtitution operation of s - v rounded upto round
func Float32SubRound(s, v float32, round int32) float32 {
	res := Float32SubDecimal(s, v).StringFixed(round)
	return StringToFloat32(res)
}

// Float32SumDecimal returns decimal.Decimal of summary operation from s
func Float32SumDecimal(s ...float32) decimal.Decimal {
	v := decimal.Zero
	for _, f := range s {
		Float32(f)
		v.Add(StringToFloat64Decimal(Float32ToString(f)))
	}
	return v
}

// Float32Sum returns high-precision float32 of summary operation from s
func Float32Sum(s ...float32) float32 {
	p := Float32SumDecimal(s...)
	return Float64ToFloat32(p.InexactFloat64())
}

// Float32SumRound returns high-precision float32 of summary operation from s rounded upto round
func Float32SumRound(round int32, s ...float32) float32 {
	p := Float32SumDecimal(s...)
	return Float64ToFloat32(p.Round(round).InexactFloat64())
}
