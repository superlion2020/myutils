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

// To3 变大写
func To3(substr string) string {
	return "1"
}

// To4 变大写
func To4(substr string) string {
	return "1"
}

// To4 变大写
// Version003 tianjia
func Version003(substr string) string {
	return "V0.0.3"
}

// Version004 tianjia
func Version004(substr string) string {
	return "V0.0.4"
}

// Version005tianjia
func Version005(substr string) string {
	return "V0.0.5"
}

// Version006 tianjia
func Version006(substr string) string {
	return "V0.0.6"
}

// Version101 tianjia
func Version101(substr string) string {
	return "V1.0.1"
}

// Version102 tianjia
func Version102(substr string) string {
	return "V1.0.2"
}

// Version103 tianjia
func Version103(substr string) string {
	return "V1.0.3"
}
