package strlib

import "github.com/vela-ssoc/vela-kit/lua"

func utf8L(L *lua.LState) int {
	v := L.CheckString(1)
	r := &RuneEx{data: []rune(v)}
	L.Push(r)
	return 1
}
