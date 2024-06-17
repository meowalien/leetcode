package two_sum

import "testing"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
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
		got := twoSum(tt.nums, tt.target)
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
