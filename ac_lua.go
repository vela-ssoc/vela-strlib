package strlib

import (
	"fmt"
	strutil "github.com/vela-ssoc/vela-kit/auxlib"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/reflectx"
)

func (m *Matcher) String() string                         { return fmt.Sprintf("%p", m) }
func (m *Matcher) Type() lua.LValueType                   { return lua.LTObject }
func (m *Matcher) AssertFloat64() (float64, bool)         { return 0, false }
func (m *Matcher) AssertString() (string, bool)           { return "", false }
func (m *Matcher) AssertFunction() (*lua.LFunction, bool) { return m.ToFunc(), true }
func (m *Matcher) Peek() lua.LValue                       { return m }

func (m *Matcher) ToFunc() *lua.LFunction {
	return lua.NewFunction(m.MatchL)
}

func (m *Matcher) MatchL(L *lua.LState) int {
	data := L.CheckString(1)

	handle := func(ret interface{}, size int) int {
		if size == 0 {
			L.Push(lua.LFalse)
			return 1
		} else {
			L.Push(lua.LTrue)
			L.Push(reflectx.ToLValue(ret, L))
			return 2
		}
	}

	if !m.extract {
		ret := m.Match(strutil.S2B(data))
		return handle(ret, len(ret))
	}

	ret := m.MAE(strutil.S2B(data))
	return handle(ret, len(ret))
}

func acL(L *lua.LState) int {
	tab := L.CheckTable(1)
	extract := L.IsTrue(2)

	patterns := tab.Strings()
	matcher := NewAcMatcher(patterns)

	if extract {
		matcher.extract = true
	}

	L.Push(matcher)
	return 1
}
