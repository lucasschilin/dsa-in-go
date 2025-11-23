package leetcode

// Given a string s, find the length of the longest substring without duplicate characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:
// - 0 <= s.length <= 5 * 104
// - s consists of English letters, digits, symbols and spaces.

func LengthOfLogestSubstring(s string) map[string]func() any {
	fns := map[string]func() any{
		"Resolution": func() any { return lengthOfLongestSubstring1(s) },
	}

	return fns
}

func lengthOfLongestSubstring1(s string) int {
	l := 0
	lastCharAppearance := make(map[rune]int)

	maxLength := 0

	for r, char := range s {

		if pos, exists := lastCharAppearance[char]; exists && pos >= l {
			l = pos + 1
		}
		lastCharAppearance[char] = r

		currentLen := r - l + 1
		if currentLen > maxLength {
			maxLength = currentLen
		}
	}

	return maxLength
}
