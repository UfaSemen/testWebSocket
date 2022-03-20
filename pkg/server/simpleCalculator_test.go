package server

import "testing"

func TestSimpleCalculator_Sum(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		sc   SimpleCalculator
		args args
		want int
	}{
		{
			name: "test on success",
			sc:   SimpleCalculator{},
			args: args{
				n1: 27,
				n2: 31,
			},
			want: 58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.Sum(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("SimpleCalculator.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCalculator_Product(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		sc   SimpleCalculator
		args args
		want int
	}{
		{
			name: "test on success",
			sc:   SimpleCalculator{},
			args: args{
				n1: 27,
				n2: 31,
			},
			want: 837,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sc.Product(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("SimpleCalculator.Product() = %v, want %v", got, tt.want)
			}
		})
	}
}
