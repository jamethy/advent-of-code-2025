package advent09

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		wantPart1 any
		wantPart2 any
	}{
		{
			name:      "sample",
			wantPart1: 50,
			wantPart2: 24,
		},
		{
			name:      "input",
			wantPart1: 4_738_108_384,
			wantPart2: 1513792010,
		},
	}
	for _, tt := range tests {
		gotPart1, gotPart2 := Solution(tt.name + ".txt")
		t.Run(tt.name+"-part1", func(t *testing.T) {
			if tt.wantPart1 != nil && !reflect.DeepEqual(gotPart1, tt.wantPart1) {
				t.Errorf("Solution() gotPart1 = %v, want %v", gotPart1, tt.wantPart1)
			}
		})
		t.Run(tt.name+"-part2", func(t *testing.T) {
			if tt.wantPart2 != nil && !reflect.DeepEqual(gotPart2, tt.wantPart2) {
				t.Errorf("Solution() gotPart2 = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}

func Test_isRightOfLine(t *testing.T) {
	type args struct {
		a    Point
		line []Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "vertical-down-right",
			args: args{
				a: Point{x: 8, y: 8},
				line: []Point{
					{x: 9, y: 3},
					{x: 9, y: 11},
				},
			},
			want: true,
		},
		{
			name: "vertical-up-right",
			args: args{
				a: Point{x: 10, y: 10},
				line: []Point{
					{x: 9, y: 11},
					{x: 9, y: 3},
				},
			},
			want: true,
		},
		{
			name: "vertical-down-left",
			args: args{
				a: Point{x: 10, y: 10},
				line: []Point{
					{x: 9, y: 3},
					{x: 9, y: 11},
				},
			},
			want: false,
		},
		{
			name: "vertical-up-left",
			args: args{
				a: Point{x: 8, y: 10},
				line: []Point{
					{x: 9, y: 11},
					{x: 9, y: 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRightOfLine(tt.args.a, tt.args.line); got != tt.want {
				t.Errorf("isRightOfLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
