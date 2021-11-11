package chineseWord

import (
	"strings"
	"testing"
)

// 每个中文三个字节，转成[]rune做长度处理，(若不转，截断出现乱码)
func TestChinaWordLenAndCut(t *testing.T) {
	s := "一二三"
	max := 1
	runeS := []rune(s)
	if len(runeS) > max {
		runeS = runeS[:max]
	}
	t.Log(len(runeS), 64 * 1024 / 3, string(runeS))
}

// 乱码 output: unknown code
func TestChinaWordLenAndCutErrWay(t *testing.T) {
	s := "一二三"
	max := 1
	if len(s) > max {
		s = s[:max]
	}
	t.Log(len(s), 64 * 1024 / 3, string(s))
}

func TestSplitWord_MatchStr4(t *testing.T) {
	const originStr = "投诉你啊，你他妈什么态度，我柔柔弱弱的我投诉你"
	r := strings.Index(originStr, "投诉")
	r2 := strings.Index(originStr, "柔柔弱弱")
	t.Logf("%d, %d\n", r2 - r, len("柔柔弱弱")) // result: 42 expect: 14(num 14 - num 0)
}

func TestUnicodeHasPrefix(t *testing.T) {
	s1 := "大佬你做什么的"
	prefix := "大佬你"
	r := UnicodeHasPrefix([]rune(s1), []rune(prefix))
	t.Logf("result: %t\n", r)
}