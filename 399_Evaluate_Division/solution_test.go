package Evaluate_Division

import (
	"reflect"
	"testing"
)

var tests = []struct {
	equations [][]string
	values    []float64
	queries   [][]string
	expected  []float64
}{
	{
		equations: [][]string{{"a", "b"}, {"b", "c"}},
		values:    []float64{2.0, 3.0},
		queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
		expected:  []float64{6.0, 0.5, -1.0, 1.0, -1.0},
	},
	{
		equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
		values:    []float64{1.5, 2.5, 5.0},
		queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
		expected:  []float64{3.75, 0.4, 5.0, 0.2},
	},
	{
		equations: [][]string{{"a", "b"}},
		values:    []float64{0.5},
		queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
		expected:  []float64{0.5, 2.0, -1.0, -1.0},
	},
	{
		equations: [][]string{{"a", "b"}, {"c", "d"}},
		values:    []float64{1.0, 1.0},
		queries:   [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}},
		expected:  []float64{-1.0, -1.0, 1.0, 1.0},
	},
}

func TestCalcEquation(t *testing.T) {
	for i, tc := range tests {
		result := calcEquation(tc.equations, tc.values, tc.queries)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed:\ngot:      %v\nexpected: %v", i, result, tc.expected)
		}
	}
}

func TestCalcEquation1(t *testing.T) {
	for i, tc := range tests {
		result := calcEquation1(tc.equations, tc.values, tc.queries)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed:\ngot:      %v\nexpected: %v", i, result, tc.expected)
		}
	}
}
