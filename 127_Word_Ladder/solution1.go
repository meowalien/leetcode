package Word_Ladder

/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.

Constraints:

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
All the words in wordList are unique.
*/
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	// 若 endWord 不在字典中，無法轉換
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	beginSet := map[string]struct{}{beginWord: {}}
	endSet := map[string]struct{}{endWord: {}}
	visited := make(map[string]struct{})
	wordLen := len(beginWord)
	level := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		// always expand the smaller set
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		nextSet := make(map[string]struct{})
		for word := range beginSet {
			wordBytes := []byte(word)
			for i := 0; i < wordLen; i++ {
				original := wordBytes[i]
				for c := byte('a'); c <= 'z'; c++ {
					if c == original {
						continue
					}
					wordBytes[i] = c
					newWord := string(wordBytes)
					if _, ok := endSet[newWord]; ok {
						return level + 1
					}
					if _, ok := wordSet[newWord]; ok {
						if _, seen := visited[newWord]; !seen {
							visited[newWord] = struct{}{}
							nextSet[newWord] = struct{}{}
						}
					}
				}
				wordBytes[i] = original
			}
		}
		beginSet = nextSet
		level++
	}
	return 0
}
