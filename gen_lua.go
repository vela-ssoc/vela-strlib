package strlib

import "github.com/vela-ssoc/vela-kit/lua"

func (s *StrGen) String() string                         { return s.raw }
func (s *StrGen) Type() lua.LValueType                   { return lua.LTObject }
func (s *StrGen) AssertFloat64() (float64, bool)         { return 0, false }
func (s *StrGen) AssertString() (string, bool)           { return "", false }
func (s *StrGen) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (s *StrGen) Peek() lua.LValue                       { return s }

func (s *StrGen) genL(L *lua.LState) int {
	L.Push(lua.S2L(s.gen()))
	return 1
}

func (s *StrGen) ret(L *lua.LState) int {
	L.Push(s)
	return 1
}

func (s *StrGen) fileL(L *lua.LState) int {
	filename := L.CheckFile(1)
	sub := L.CheckString(2)

	s.fnc = append(s.fnc, NewTrimAcFile(filename, sub))
	return s.ret(L)
}

func (s *StrGen) dateL(L *lua.LState) int {
	sub := L.CheckString(1)
	fm := L.Get(2)

	s.fnc = append(s.fnc, NewTrimDate(sub, fm.String()))
	return s.ret(L)
}

func (s *StrGen) numL(L *lua.LState) int {
	s.fnc = append(s.fnc, NewTrimN())
	return s.ret(L)
}

func (s *StrGen) graphicL(L *lua.LState) int {
	flag := L.IsTrue(1)
	s.fnc = append(s.fnc, NewTrimGraphic(flag))
	return s.ret(L)
}

func (s *StrGen) spaceL(L *lua.LState) int {
	s.fnc = append(s.fnc, NewTrimSpace())
	return s.ret(L)
}

func (s *StrGen) regexL(L *lua.LState) int {
	regex := L.CheckString(1)
	sub := L.CheckString(2)
	s.fnc = append(s.fnc, NewTrimRegex(regex, sub))
	return s.ret(L)
}

func (s *StrGen) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "gen":
		return L.NewFunction(s.genL)
	case "date":
		return L.NewFunction(s.dateL)
	case "num":
		return L.NewFunction(s.numL)
	case "graphic":
		return L.NewFunction(s.graphicL)
	case "space":
		return L.NewFunction(s.spaceL)
	case "regex":
		return L.NewFunction(s.regexL)
	case "file":
		return L.NewFunction(s.fileL)
	}

	return lua.LNil
}

func generalizationL(L *lua.LState) int {
	str := L.CheckString(1)
	gen := &StrGen{raw: str}
	L.Push(gen)
	return 1
}
