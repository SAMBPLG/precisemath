package precisemath

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestPercentage(t *testing.T) {
	type args struct {
		of            float64
		from          float64
		decimalLength int32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "RoundNumber",
			args: args{
				of:            1,
				from:          10,
				decimalLength: 0,
			},
			want: 10,
		},
		{
			name: "FloatNumber",
			args: args{
				of:            25541546.52,
				from:          459251427.458751,
				decimalLength: 2,
			},
			want: 5.5615606,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Percentage(tt.args.of, tt.args.from); got != tt.want {
				t.Errorf("Percentage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentageRound(t *testing.T) {
	type args struct {
		of            float64
		from          float64
		decimalLength int32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "RoundNumber",
			args: args{
				of:            1,
				from:          10,
				decimalLength: 0,
			},
			want: 10,
		},
		{
			name: "FloatNumber",
			args: args{
				of:            25541546.52,
				from:          459251427.458751,
				decimalLength: 2,
			},
			want: 5.56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageRound(tt.args.of, tt.args.from, tt.args.decimalLength); got != tt.want {
				t.Errorf("PercentageRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentageFromString(t *testing.T) {
	type args struct {
		of            string
		from          string
		decimalLength int32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "NoExponent",
			args: args{
				of:            "25541546.52",
				from:          "459251427.458751",
				decimalLength: 2,
			},
			want: 5.5615606,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageFromString(tt.args.of, tt.args.from); got != tt.want {
				t.Errorf("PercentageFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentageFromStringDecimal(t *testing.T) {
	type args struct {
		s string
		v string
	}
	tests := []struct {
		name string
		args args
		want decimal.Decimal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageFromStringDecimal(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PercentageFromStringDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentageFromStringRound(t *testing.T) {
	type args struct {
		of    string
		from  string
		round int32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageFromStringRound(tt.args.of, tt.args.from, tt.args.round); got != tt.want {
				t.Errorf("PercentageFromStringRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmountOfDecimal(t *testing.T) {
	type args struct {
		s float64
		v float64
	}
	tests := []struct {
		name string
		args args
		want decimal.Decimal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AmountOfDecimal(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AmountOfDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmountOf(t *testing.T) {
	type args struct {
		s float64
		v float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AmountOf(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("AmountOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmountOfRound(t *testing.T) {
	type args struct {
		s     float64
		v     float64
		round int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AmountOfRound(tt.args.s, tt.args.v, tt.args.round); got != tt.want {
				t.Errorf("AmountOfRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
