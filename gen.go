package strlib

type FunctionGen func(string) string

type StrGen struct {
	raw string
	fnc []FunctionGen
}

func (s *StrGen) gen() string {
	str := s.raw
	n := len(s.fnc)
	if n == 0 {
		return str
	}

	for i := 0; i < n; i++ {
		str = s.fnc[i](str)
	}

	return str
}
