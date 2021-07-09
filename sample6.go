// Problem statement 1
//
// Given a string, find the length of the longest substring without repeating characters.
// Examples:
// - Given "abcabcbb", the answer is "abc", which the length is 3.
// - Given "bbbbb", the answer is "b", with the length of 1.
// - Given "pwwkew", the answer is "wke", with the length of 3.
// - Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

func main() {
	fmt.Println("abcabcbb -> ", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("bbbbbbbb -> ", lengthOfLongestSubstring("bbbbbbbb"))
	fmt.Println("pwwkew -> ", lengthOfLongestSubstring("pwwkew"))
}
