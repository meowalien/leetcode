package Snakes_and_Ladders

import (
	"testing"
)

var tests = []struct {
	board [][]int
	want  int
}{
	{
		board: [][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1},
		},
		want: 4,
	},
	{
		board: [][]int{
			{-1, -1},
			{-1, 3},
		},
		want: 1,
	},
	{
		board: [][]int{
			{-1, -1, -1},
			{-1, 9, 8},
			{-1, 8, 9},
		},
		want: 1,
	},
	{
		board: [][]int{
			{-1, -1, -1},
			{-1, -1, -1},
			{-1, -1, -1},
		},
		want: 2,
	},
}

func TestSnakesAndLadders(t *testing.T) {
	for _, tt := range tests {
		got := snakesAndLadders(tt.board)
		if got != tt.want {
			t.Errorf("snakesAndLadders() = %d, want %d", got, tt.want)
		}
	}
}
