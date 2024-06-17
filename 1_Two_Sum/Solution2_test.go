package two_sum

import "testing"

func twoSum1(nums []int, target int) []int {
	luigi := map[int]int{}

	for i, el := range nums {
		if _, ok := luigi[el]; ok {
			return []int{luigi[el], i}
		}
		luigi[target-el] = i
	}
	return nil
}

func TestTwoSum1(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		got := twoSum1(tt.nums, tt.target)
		if len(got) != len(tt.want) {
			t.Errorf("twoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
			continue
		}
		for i, v := range got {
			if v != tt.want[i] {
				t.Errorf("twoSum(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
				break
			}
		}
	}
}
