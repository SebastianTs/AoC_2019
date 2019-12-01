package main

import (
	"testing"
)

func Test_calcFuelRequired(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{12}, 2},
		{"example2", args{14}, 2},
		{"example3", args{1969}, 654},
		{"example4", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuelRequired(tt.args.mass); got != tt.want {
				t.Errorf("calcFuelRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFuelRequiredPartTwo(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{14}, 2},
		{"example2", args{1969}, 966},
		{"example3", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuelRequiredPartTwo(tt.args.mass); got != tt.want {
				t.Errorf("calcFuelRequiredPartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
