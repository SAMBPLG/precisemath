package precisemath

import (
	"math"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestFloat64(t *testing.T) {
	tests := []float64{1.123456789, math.Pi, -0.0000001}
	for _, v := range tests {
		got := Float64(v)
		if math.Abs(got-v) > 1e-12 {
			t.Errorf("Float64(%v) = %v; want ~%v", v, got, v)
		}
	}
}

func TestFloat64ToString(t *testing.T) {
	tests := map[float64]string{
		1.5:            "1.5",
		123456789.123:  "1.23456789123e+08",
		0.000000012345: "1.2345e-08", // default formatting,
		2.721625315:    "2.721625315",
	}
	for input, expected := range tests {
		got := Float64ToString(input)
		if got != expected {
			t.Errorf("Float64ToString(%v) = %s; want %s", input, got, expected)
		}
	}
}

func TestFloat64ToStringScientific(t *testing.T) {
	tests := map[float64]string{
		1.5:            "1.5e+00",
		123456789.123:  "1.23456789123e+08",
		0.000000012345: "1.2345e-08",
		2.721625315:    "2.721625315e+00",
	}
	for input, expected := range tests {
		got := Float64ToStringScientific(input)
		if got != expected {
			t.Errorf("Float64ToStringScientific(%v) = %s; want %s", input, got, expected)
		}
	}
}

func TestFloat64Round(t *testing.T) {
	got := Float64Round(1.23456789, 4)
	expected := 1.2346
	if math.Abs(got-expected) > 1e-6 {
		t.Errorf("Float64Round(1.23456789, 4) = %v; want %v", got, expected)
	}
}

func TestFloat64ToFloat32(t *testing.T) {
	input := 1.23456789
	got := Float64ToFloat32(input)
	expected := float32(input)
	if math.Abs(float64(got)-float64(expected)) > 1e-6 {
		t.Errorf("Float64ToFloat32(%v) = %v; want %v", input, got, expected)
	}
}

func TestFloat64DivDecimal(t *testing.T) {
	got := Float64DivDecimal(10.0, 4.0)
	expected := decimal.NewFromFloat(2.5)
	if !got.Equal(expected) {
		t.Errorf("Float64DivDecimal(10.0, 4.0) = %v; want %v", got, expected)
	}
	zeroDiv := Float64DivDecimal(10.0, 0.0)
	if !zeroDiv.Equal(decimal.Zero) {
		t.Errorf("Float64DivDecimal(10.0, 0.0) = %v; want 0", zeroDiv)
	}
}

func TestFloat64SubDecimal(t *testing.T) {
	got := Float64SubDecimal(5.5, 2.2)
	expected := decimal.NewFromFloat(3.3)
	if !got.Equal(expected) {
		t.Errorf("Float64SubDecimal(5.5, 2.2) = %v; want %v", got, expected)
	}
}

func TestFloat64Sub(t *testing.T) {
	got := Float64Sub(5.5, 2.2)
	expected := 3.3
	if math.Abs(got-expected) > 1e-9 {
		t.Errorf("Float64Sub(5.5, 2.2) = %v; want %v", got, expected)
	}
}

func TestFloat64SubRound(t *testing.T) {
	got := Float64SubRound(5.55555, 2.22222, 2)
	expected := 3.33
	if math.Abs(got-expected) > 1e-2 {
		t.Errorf("Float64SubRound(5.55555, 2.22222, 2) = %v; want %v", got, expected)
	}
}

func TestFloat64SumDecimal(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
		want decimal.Decimal
	}{
		{
			name: "RoundNumber",
			args: args{s: []float64{1.5, 1.5, 1.5}},
			want: decimal.NewFromFloat(4.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64SumDecimal(tt.args.s...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64SumDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Sum(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "RoundNumber",
			args: args{s: []float64{1.5, 1.5, 1.5}},
			want: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Sum(tt.args.s...); got != tt.want {
				t.Errorf("Float64Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64SumRound(t *testing.T) {
	type args struct {
		round int32
		s     []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "RoundNumber",
			args: args{round: 2, s: []float64{1, 1, 1}},
			want: 3,
		},
		{
			name: "RoundNumber",
			args: args{round: 2, s: []float64{1.5, 1.5, 1.5}},
			want: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64SumRound(tt.args.round, tt.args.s...); got != tt.want {
				t.Errorf("Float64SumRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
