package client

import (
	"reflect"
	"testing"
)

func TestNewRandomGenerator(t *testing.T) {
	tests := []struct {
		name string
		want RandomGenerator
	}{
		{
			name: "test on success",
			want: RandomGenerator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomGenerator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomGenerator_Generate(t *testing.T) {
	tests := []struct {
		name string
		rg   RandomGenerator
	}{
		{
			name: "test on success",
			rg:   RandomGenerator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rg.Generate(); got > 100 {
				t.Errorf("RandomGenerator.Generate() = %v", got)
			}
		})
	}
}
