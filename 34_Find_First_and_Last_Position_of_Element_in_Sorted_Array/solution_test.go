package Find_First_and_Last_Position_of_Element_in_Sorted_Array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2, 2}},
		{[]int{2, 2, 2, 2, 2}, 2, []int{0, 4}},
		{[]int{1, 3, 3, 3, 5, 5, 9}, 5, []int{4, 5}},
	}

	for _, test := range tests {
		result := searchRange(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums = %v, target = %d, expected %v, but got %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
