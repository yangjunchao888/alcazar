package chineseWord

import (
	"strings"
)

func UnicodeIndex(str string, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func UnicodeHasPrefix(s, prefix []rune) bool {
	if len(s) < len(prefix) {
		return false
	}

	for i:=0; i<len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}

	return true
}
