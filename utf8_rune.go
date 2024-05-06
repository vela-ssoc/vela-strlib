package strlib

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

type RuneEx struct {
	data []rune
}

func (r *RuneEx) String() string                         { return string(r.data) }
func (r *RuneEx) Type() lua.LValueType                   { return lua.LTObject }
func (r *RuneEx) AssertFloat64() (float64, bool)         { return 0, false }
func (r *RuneEx) AssertString() (string, bool)           { return r.String(), true }
func (r *RuneEx) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (r *RuneEx) Peek() lua.LValue                       { return r }

func (r *RuneEx) pipeL(L *lua.LState) int {
	chain := pipe.NewByLua(L, pipe.Env(xEnv))

	for _, v := range r.data {
		chain.Do(lua.LInt(v), L, func(x error) {
			//todo
		})
	}

	return 0
}

//r.trim("10-100" , "100" , "200")

func (r *RuneEx) trimL(L *lua.LState) int {
	if len(r.data) == 0 {
		return 0
	}

	s := L.IsInt(1)
	e := L.IsInt(2)

	inversion := L.IsTrue(3)

	cmp := func(v rune) bool {
		if v >= int32(s) && v <= int32(e) {
			return true
		}

		return false
	}

	buff := make([]rune, len(r.data))
	offset := 0
	for _, item := range r.data {
		flag := cmp(item)

		if flag && !inversion {
			buff[offset] = item
			offset++
			continue
		}

		if !flag && inversion {
			buff[offset] = item
			offset++
			continue
		}
	}

	r.data = buff[:offset]
	return 0
}

func (r *RuneEx) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "pipe":
		return lua.NewFunction(r.pipeL)
	case "size":
		return lua.LInt(len(r.data))
	case "trim":
		return lua.NewFunction(r.trimL)
	case "text":
		return lua.S2L(string(r.data))
	}
	return lua.LNil
}
