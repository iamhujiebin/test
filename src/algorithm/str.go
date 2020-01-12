package algorithm

import (
	"bytes"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	t := []rune(s)
	var str string
	p, q := 0, 0
	num := 1
	max := num
	str = string(t[0])
	for q+1 < len(t) {
		q = q + 1
		if strings.Contains(str, string(t[q])) {
			p++
			q = p
			if num > max {
				max = num
			}
			num = 1
			str = string(t[p])
		} else {
			str = str + string(t[q])
			num++
		}
	}
	if num > max {
		max = num
	}
	return max
}

func LengthOfLongestSubstring2(s string) int {
	if len(s) <= 0 {
		return 0
	}
	runes := []rune(s)
	var str string
	str = string(runes[0])
	maxLen := 1
	for _, v := range []rune(s) {
		if index := strings.Index(str, string(v)); index != -1 {
			str = str[index+1:]
			str = str + string(v)
		} else {
			str = str + string(v)
			if len([]rune(str)) > maxLen {
				maxLen = len([]rune(str))
			}
		}
	}
	return maxLen
}

func LengthOfLongestSubstringTarget(s string) int {
	ss := make([]byte, 0, 0)
	maxLen := 0
	for _, v := range []byte(s) {
		if index := bytes.IndexByte(ss, v); index != -1 {
			ss = ss[index+1:]
			ss = append(ss, v)
		} else {
			ss = append(ss, v)
			if maxLen < len(ss) {
				maxLen = len(ss)
			}
		}
	}
	return maxLen
}

func LengthOfLongestSubStringTarget2(s string) int {
	flag := [256]int{}
	beg := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if flag[s[i]] > 0 && flag[s[i]] > beg {
			beg = flag[s[i]]
		}
		flag[s[i]] = i + 1
		max = maxnum(max, i-beg+1)
	}
	return max

}

func maxnum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LongestPalindrome(s string) string {
	var index int
	var maxLen int
	for i := range s {
		len1 := expand(s, i, i)
		len2 := expand(s, i, i+1)
		m := max(len1, len2)
		if m > maxLen {
			index = i
			maxLen = m
		}
	}
	if maxLen < 1 {
		return ""
	}
	if maxLen%2 == 0 { // "abba" 的回文
		start := index - (maxLen/2 - 1)
		end := index + maxLen/2
		return s[start : end+1]
	} else { // "abcba" 的回文
		start := index - maxLen/2
		end := index + maxLen/2
		return s[start : end+1]
	}
}

func expand(s string, i, j int) (strLen int) {
	L, R := i, j
	for L >= 0 && R < len(s) {
		if s[L] == s[R] {
			if len(s[L:R+1]) > strLen {
				strLen = len(s[L : R+1])
			}
			L--
			R++
		} else {
			break
		}
	}
	return
}
