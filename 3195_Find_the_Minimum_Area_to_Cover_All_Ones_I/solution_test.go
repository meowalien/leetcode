package Find_the_Minimum_Area_to_Cover_All_Ones_I

import "testing"

func TestSolution1(t *testing.T) {
	grid := [][]int{{0, 1, 0}, {1, 0, 1}}
	expected := 6
	got := minimumArea(grid)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}

func TestSolution2(t *testing.T) {
	grid := [][]int{{0, 0}, {1, 0}}
	expected := 1
	got := minimumArea(grid)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}
