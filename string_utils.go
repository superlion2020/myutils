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

// ToDa daxia
func ToDa(substr string) string {
	return strings.ToUpper(substr)
}

// ToXiao 变大写
func ToXiao(substr string) string {
	return strings.ToUpper(substr)
}

// ToXiao 变大写
func To1(substr string) string {
	return "1"
}

// To2 变大写
func To2(substr string) string {
	return "1"
}
