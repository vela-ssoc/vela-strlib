package strlib

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/vela"
)

var xEnv vela.Environment

type StrLib struct{}

func (s *StrLib) String() string                         { return "lua.strlib.export" }
func (s *StrLib) Type() lua.LValueType                   { return lua.LTObject }
func (s *StrLib) AssertFloat64() (float64, bool)         { return 0, false }
func (s *StrLib) AssertString() (string, bool)           { return "", false }
func (s *StrLib) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (s *StrLib) Peek() lua.LValue                       { return s }

func (s *StrLib) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "utf8":
		return lua.NewFunction(utf8L)
	case "similarity":
		return lua.NewFunction(similarityL)
	case "ac":
		return lua.NewFunction(acL)
	case "gen":
		return lua.NewFunction(generalizationL)
	default:
		return lua.LNil
	}
}

func WithEnv(env vela.Environment) {
	xEnv = env
	xEnv.Set("strlib", new(StrLib))
}
