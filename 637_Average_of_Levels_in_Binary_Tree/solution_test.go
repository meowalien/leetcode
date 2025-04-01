package Average_of_Levels_in_Binary_Tree

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	root     *TreeNode
	expected []float64
}{
	{
		name: "Example 1",
		root: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		expected: []float64{3.00000, 14.50000, 11.00000},
	},
	{
		name: "Example 2",
		root: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 9,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{Val: 20},
		},
		expected: []float64{3.00000, 14.50000, 11.00000},
	},
	{
		name:     "Single Node",
		root:     &TreeNode{Val: 1},
		expected: []float64{1.00000},
	},
}

func TestAverageOfLevels(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := averageOfLevels(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}

func TestAverageOfLevels1(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := averageOfLevels1(tc.root)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expected, result)
			}
		})
	}
}
