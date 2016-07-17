package strings

import "unicode/utf8"

func RepeatChar(c rune, count int) string {
	s := make([]rune, count)
	for i := range s {
		s[i] = c
	}
	return string(s)
}

func StringLen(s string) (bytes int, runes int) {
	b    := []byte(s)
	bytes = len(b)
	runes = utf8.RuneCount(b)
	return
}

func StringReplace(s string, pos int, r string) string {
	rs  := []rune(s)
	rr  := []rune(r)
	copy(rs[pos:pos+len(rr)], rr)
	return string(rs)
}


