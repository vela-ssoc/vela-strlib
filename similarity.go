package strlib

import (
	"fmt"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/strutil"
	"github.com/vela-ssoc/vela-kit/strutil/similarity"
)

type Similarity struct {
	a string
	b string
}

func (s Similarity) String() string                         { return fmt.Sprintf("%p", &s) }
func (s Similarity) Type() lua.LValueType                   { return lua.LTObject }
func (s Similarity) AssertFloat64() (float64, bool)         { return 0, false }
func (s Similarity) AssertString() (string, bool)           { return "", false }
func (s Similarity) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (s Similarity) Peek() lua.LValue                       { return s }

func (s Similarity) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "levenshtein":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewLevenshtein()))

	case "hamming":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewHamming()))

	case "jaro":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewJaro()))

	case "jacc":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewJaccard()))

	case "jaro_winkler":

		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewJaroWinkler()))
	case "overlap":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewOverlapCoefficient()))

	case "smithwatermangotoh":

		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewSmithWatermanGotoh()))

	case "sorensendice":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewSorensenDice()))

	case "prop":
		return lua.LNumber(strutil.Similarity(s.a, s.b, similarity.NewProportion()))

	}

	return lua.LNil

}

func similarityL(L *lua.LState) int {
	a := L.ToString(1)
	b := L.ToString(2)
	L.Push(Similarity{a: a, b: b})
	return 1

}
