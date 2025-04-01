package Number_of_1_Bits

import "testing"

var tests = []struct {
	input    int
	expected int
}{
	{11, 3},          // binary: 1011
	{128, 1},         // binary: 10000000
	{2147483645, 30}, // binary: 1111111111111111111111111111101
	{0, 0},           // binary: 0
	{4294967295, 32}, // binary: all 1s (uint32 max value)
}

func TestHammingWeight(t *testing.T) {
	for _, test := range tests {
		result := hammingWeight(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d but got %d", test.input, test.expected, result)
		}
	}
}

func TestHammingWeight1(t *testing.T) {
	for _, test := range tests {
		result := hammingWeight1(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
