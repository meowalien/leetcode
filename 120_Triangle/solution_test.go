package _20_Triangle

import "testing"

var fc = []func(triangle [][]int) int{minimumTotal, minimumTotal1}

func TestSolution1(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	correct := 11
	for _, f := range fc {
		result := f(triangle)
		if result != correct {
			t.Fatalf("incorrect result: %v", result)
		}
	}
}

func TestSolution2(t *testing.T) {
	triangle := [][]int{{-10}}
	correct := -10
	for _, f := range fc {
		result := f(triangle)
		if result != correct {
			t.Fatalf("incorrect result: %v", result)
		}
	}
}

func TestSolution3(t *testing.T) {
	triangle := [][]int{{1}, {2, 3}}
	correct := 3
	for _, f := range fc {
		result := f(triangle)
		if result != correct {
			t.Fatalf("incorrect result: %v", result)
		}
	}
}
