package myutils

import (
	"strings"
)

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ContainsIgnoreCase 忽略大小写检查字符串是否包含子串
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// 变大写
func RemoveFirstOne(substr string) string {
	return strings.ToUpper(substr)
}
