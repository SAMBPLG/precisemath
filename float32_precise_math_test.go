package precisemath

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestFloat32(t *testing.T) {
	type args struct {
		of float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "ShouldEqual",
			args: args{
				of: 1.00123,
			},
			want: 1.00123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.of); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32ToString(t *testing.T) {
	type args struct {
		of float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ShouldEqual",
			args: args{
				of: 1.00123,
			},
			want: "1.00123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32ToString(tt.args.of); got != tt.want {
				t.Errorf("Float32ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32ToStringScientific(t *testing.T) {
	type args struct {
		of float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ShouldEqual",
			args: args{
				of: 1.00123,
			},
			want: "1.00123e+00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32ToStringScientific(tt.args.of); got != tt.want {
				t.Errorf("Float32ToStringScientific() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Round(t *testing.T) {
	type args struct {
		of    float32
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
			if got := Float32Round(tt.args.of, tt.args.round); got != tt.want {
				t.Errorf("Float32Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32ToFloat64(t *testing.T) {
	type args struct {
		of float32
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
			if got := Float32ToFloat64(tt.args.of); got != tt.want {
				t.Errorf("Float32ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32DivDecimal(t *testing.T) {
	type args struct {
		s float32
		v float32
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
			if got := Float32DivDecimal(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32DivDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32SubDecimal(t *testing.T) {
	type args struct {
		s float32
		v float32
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
			if got := Float32SubDecimal(tt.args.s, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32SubDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Sub(t *testing.T) {
	type args struct {
		s float32
		v float32
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
			if got := Float32Sub(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("Float32Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32SubRound(t *testing.T) {
	type args struct {
		s     float32
		v     float32
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
			if got := Float32SubRound(tt.args.s, tt.args.v, tt.args.round); got != tt.want {
				t.Errorf("Float32SubRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32SumDecimal(t *testing.T) {
	type args struct {
		s []float32
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
			if got := Float32SumDecimal(tt.args.s...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32SumDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Sum(t *testing.T) {
	type args struct {
		s []float32
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
			if got := Float32Sum(tt.args.s...); got != tt.want {
				t.Errorf("Float32Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32SumRound(t *testing.T) {
	type args struct {
		round int32
		s     []float32
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
			if got := Float32SumRound(tt.args.round, tt.args.s...); got != tt.want {
				t.Errorf("Float32SumRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
