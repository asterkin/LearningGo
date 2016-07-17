package strings

import "testing"

func TestRepeatChar(t *testing.T) {
	repeatTests := []struct {
		count    int
		expected string
	}{
		{0, ""},
		{1, "A"},
		{5, "AAAAA"},
	}

	for _,c:=range repeatTests {
		actual := RepeatChar('A', c.count)
		if actual != c.expected {
			t.Errorf("RepeatChar(%d) expected '%s' actual '%s'", c.count, c.expected, actual)
		}
	}
}

func TestStringLen(t *testing.T) {
	stringLenTests := [] struct {
		s     string
 		bytes int
		runes int
	} {
		{"",                       0, 0},
		{"asSASA ddd dsjkdsjs dk", 22, 22},
	}
	for _,c := range stringLenTests {
		bytes, runes := StringLen(c.s)
		if bytes != c.bytes  || runes != c.runes {
			t.Errorf("StringLen(%s) expected (%d, %d) actual (%d, %d)",
				c.s,
				c.bytes,
				c.runes,
				bytes, runes)
		}
	}
}

func TestStringReplace(t *testing.T) {
	stringReplaceTests := [] struct {
		s        string
		pos      int
		r        string
		expected string
	}{
		{"abcdefg",                1, "BC",  "aBCdefg"},
		{"asSASA ddd dsjkdsjs dk", 4, "abc", "asSAabcddd dsjkdsjs dk"},
	}
	for _,c := range stringReplaceTests {
		actual := StringReplace(c.s, c.pos, c.r)
		if actual != c.expected {
			t.Errorf("StringRepleace(%s, %d, %s) expected '%s' actual '%s'",
				c.s, c.pos, c.r, c.expected, actual)
		}
	}
}

func TestStringReverse(t *testing.T) {
	stringReverseTests := [] struct {
		s        string
		expected string
	}{
		{"raboof", "foobar"},
		{"ducktyping", "gnipytkcud"},
	}

	for _,c := range stringReverseTests {
		actual := StringReverse(c.s)
		if actual != c.expected {
			t.Errorf("StringReverse(%s) expected '%s' actual '%s'", c.s, c.expected, actual)
		}
	}
}

func StringReverse(s string) string {
	rs := []rune(s)
	for i,j :=0, len(rs)-1; i<j; i,j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}


