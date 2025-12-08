package advent08

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name       string
		iterations int
		wantPart1  any
		wantPart2  any
	}{
		{
			name:       "sample",
			iterations: 10,
			wantPart1:  40,
			wantPart2:  25272,
		},
		{
			name:       "input",
			iterations: 1000,
			wantPart1:  127551,
			wantPart2:  2347225200,
		},
	}
	for _, tt := range tests {
		gotPart1, gotPart2 := Solution(tt.name+".txt", tt.iterations)
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
