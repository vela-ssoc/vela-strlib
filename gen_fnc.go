package strlib

import (
	strutil "github.com/vela-ssoc/vela-kit/auxlib"
	"regexp"
	"strings"
	"unicode"
)

const (
	N0 = '0'
	N9 = '9'
	NA = 'a'
	NZ = 'z'
)

func NewTrimN() FunctionGen {
	return func(str string) string {
		u := []rune(str)
		n := len(u)
		for i := 0; i < n; i++ {
			if u[i] >= N0 && u[i] <= N9 {
				u[i] = 'N'
			}
		}
		return string(u)
	}
}

func NewTrimGraphic(flag bool) FunctionGen {
	return func(str string) string {
		return strings.TrimFunc(str, func(r rune) bool {
			return !unicode.IsGraphic(r)
		})
	}
}

func NewTrimSpace() FunctionGen {
	return func(s string) string {
		return strings.TrimFunc(s, func(r rune) bool {
			return unicode.IsSpace(r)
		})
	}
}

func NewTrimAcFile(filename string, ch string) FunctionGen {
	return func(s string) string {
		ac, err := NewAcFile(filename)
		if err != nil {
			return s
		}
		return strutil.B2S(ac.Replace(strutil.S2B(s), strutil.S2B(ch)))
	}
}

func NewTrimRegex(regex string, ch string) FunctionGen {
	return func(s string) string {
		re, err := regexp.Compile(regex)
		if err != nil {
			return s
		}

		return re.ReplaceAllString(s, ch)
	}
}

func NewTrimDate(f string, m string) FunctionGen {
	return func(str string) string {
		return strings.ReplaceAll(str, f, m)
	}
}
