package Count_Pairs_That_Form_a_Complete_Day_I

import "testing"

var fc = []func(hours []int) int{countCompleteDayPairs}

func testAll(t *testing.T, input []int, correct int) {
	for _, f := range fc {
		res := f(input)
		if res != correct {
			t.Errorf("expected 2, got %d", res)
		}
	}
}

func TestSolution1(t *testing.T) {
	testAll(t, []int{12, 12, 30, 24, 24}, 2)
}

func TestSolution2(t *testing.T) {
	testAll(t, []int{72, 48, 24, 3}, 3)
}
