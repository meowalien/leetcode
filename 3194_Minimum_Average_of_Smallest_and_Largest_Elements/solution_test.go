package _194_Minimum_Average_of_Smallest_and_Largest_Elements

import "testing"

func TestSolution1(t *testing.T) {
	nums := []int{7, 8, 3, 4, 15, 13, 4, 1}
	expected := 5.5
	got := minimumAverage(nums)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}

func TestSolution2(t *testing.T) {
	nums := []int{1, 9, 8, 3, 10, 5}
	expected := 5.5
	got := minimumAverage(nums)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}

func TestSolution3(t *testing.T) {
	nums := []int{1, 2, 3, 7, 8, 9}
	expected := 5.0
	got := minimumAverage(nums)
	if got != expected {
		t.Fatalf("got: %v, expected: %v", got, expected)
	}
}
