package Word_Ladder

import "testing"

var tests = []struct {
	name      string
	beginWord string
	endWord   string
	wordList  []string
	expected  int
}{
	{
		name:      "Example 1",
		beginWord: "hit",
		endWord:   "cog",
		wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
		expected:  5,
	},
	{
		name:      "Example 2 - endWord not in wordList",
		beginWord: "hit",
		endWord:   "cog",
		wordList:  []string{"hot", "dot", "dog", "lot", "log"},
		expected:  0,
	},
	{
		name:      "Single-letter words with valid transformation",
		beginWord: "a",
		endWord:   "c",
		wordList:  []string{"a", "b", "c"},
		expected:  2, // "a" -> "c"
	},
	{
		name:      "No valid transformation available",
		beginWord: "hit",
		endWord:   "xyz",
		wordList:  []string{"hot", "dot", "dog"},
		expected:  0,
	},
	{
		name:      "Lead to Gold transformation",
		beginWord: "lead",
		endWord:   "gold",
		wordList:  []string{"load", "goad", "gold"},
		expected:  4, // "lead" -> "load" -> "goad" -> "gold"
	},
}

func TestLadderLength(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if result != tt.expected {
				t.Errorf("ladderLength(%q, %q, %v) = %d; want %d",
					tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
			}
		})
	}
}

func TestLadderLength1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ladderLength1(tt.beginWord, tt.endWord, tt.wordList)
			if result != tt.expected {
				t.Errorf("ladderLength(%q, %q, %v) = %d; want %d",
					tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
			}
		})
	}
}
