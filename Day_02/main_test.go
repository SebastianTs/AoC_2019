package main

import (
	"testing"
)

func Test_processIntcode(t *testing.T) {
	type args struct {
		ns []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{[]int{1, 0, 0, 0, 99}}, []int{2, 0, 0, 0, 99}},
		{"example2", args{[]int{2, 3, 0, 3, 99}}, []int{2, 3, 0, 6, 99}},
		{"example3", args{[]int{2, 4, 4, 5, 99, 0}}, []int{2, 4, 4, 5, 99, 9801}},
		{"example4", args{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processIntcode(tt.args.ns)
			for i, v := range tt.want {
				if tt.args.ns[i] != v {
					t.Errorf("processIntcode() = %v, want %v at index %v\ngot:\t%v\nwanted:\t%v", tt.args.ns[i], v, i, tt.args.ns, tt.want)
				}
			}
		})
	}
}

func Test_findOutput(t *testing.T) {
	type args struct {
		output int
		ns     []int
	}
	tests := []struct {
		name     string
		args     args
		wantNoun int
		wantVerb int
		wantErr  bool
	}{
		{"example1",
			args{19690720,
				[]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 5, 19, 23, 2, 9, 23, 27, 1, 6, 27,
					31, 1, 31, 9, 35, 2, 35, 10, 39, 1, 5, 39, 43, 2, 43, 9, 47, 1, 5, 47, 51, 1, 51, 5, 55, 1, 55,
					9, 59, 2, 59, 13, 63, 1, 63, 9, 67, 1, 9, 67, 71, 2, 71, 10, 75, 1, 75, 6, 79, 2, 10, 79, 83, 1,
					5, 83, 87, 2, 87, 10, 91, 1, 91, 5, 95, 1, 6, 95, 99, 2, 99, 13, 103, 1, 103, 6, 107, 1, 107,
					5, 111, 2, 6, 111, 115, 1, 115, 13, 119, 1, 119, 2, 123, 1, 5, 123, 0, 99, 2, 0, 14, 0},
			},
			84,
			78,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNoun, gotVerb, err := findOutput(tt.args.output, tt.args.ns)
			if (err != nil) != tt.wantErr {
				t.Errorf("findOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNoun != tt.wantNoun {
				t.Errorf("findOutput() gotNoun = %v, want %v", gotNoun, tt.wantNoun)
			}
			if gotVerb != tt.wantVerb {
				t.Errorf("findOutput() gotVerb = %v, want %v", gotVerb, tt.wantVerb)
			}
		})
	}
}
