package precisemath

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestStringToFloat64(t *testing.T) {
	type args struct {
		of string
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
			if got := StringToFloat64(tt.args.of); got != tt.want {
				t.Errorf("StringToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat64Decimal(t *testing.T) {
	type args struct {
		of string
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
			if got := StringToFloat64Decimal(tt.args.of); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToFloat64Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat64Round(t *testing.T) {
	type args struct {
		of    string
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
			if got := StringToFloat64Round(tt.args.of, tt.args.round); got != tt.want {
				t.Errorf("StringToFloat64Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat32(t *testing.T) {
	type args struct {
		of string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "ShouldEqual_WithExponential",
			args: args{
				of: "1.00123e+00",
			},
			want: 1.00123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToFloat32(tt.args.of); got != tt.want {
				t.Errorf("StringToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToFloat32Round(t *testing.T) {
	type args struct {
		of    string
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
			if got := StringToFloat32Round(tt.args.of, tt.args.round); got != tt.want {
				t.Errorf("StringToFloat32Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
