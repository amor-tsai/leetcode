/*
3.Longest Substring Without Repeating Characters

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("pwlgkabwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var tmpString []byte
	maxLen := 0
	if 0 == len(s) {
		return 0
	}
	for i := 0; i < len(s); i++ {
		index := bytes.Index(tmpString, []byte{s[i]})
		if -1 != index {
			if index == (len(tmpString) - 1) {
				tmpString = []byte{s[i]}
			} else {
				tmpString = append(tmpString, s[i])
				tmpString = tmpString[(index + 1):]
			}
		} else {
			tmpString = append(tmpString, s[i])
		}
		if len(tmpString) > maxLen {
			maxLen = len(tmpString)
		}
		// fmt.Println(i, ":", string(tmpString))
	}
	return maxLen
}
