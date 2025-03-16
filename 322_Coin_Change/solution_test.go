package Coin_Change

import "testing"

func TestCoinChange(t *testing.T) {
	// Define test cases with coins, amount, and expected output.
	testCases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		// Example test cases from the problem description.
		//{coins: []int{1, 2, 5}, amount: 11, expected: 3},      // 11 = 5 + 5 + 1
		//{coins: []int{2}, amount: 3, expected: -1},            // cannot form amount 3
		//{coins: []int{1}, amount: 0, expected: 0},             // zero amount requires 0 coins
		//{coins: []int{1, 2, 3}, amount: 6, expected: 2},       // 6 = 3 + 3
		//{coins: []int{2, 5, 10, 1}, amount: 27, expected: 4},  // 27 = 10 + 10 + 5 + 2
		//{coins: []int{1, 2147483647}, amount: 2, expected: 2}, // 2 = 1 + 1
		{coins: []int{186, 419, 83, 408}, amount: 6249, expected: 20},
	}

	// Iterate over each test case.
	for _, tc := range testCases {
		result := coinChange(tc.coins, tc.amount)
		if result != tc.expected {
			t.Errorf("coinChange(%v, %d) = %d; expected %d", tc.coins, tc.amount, result, tc.expected)
		}
	}
}
