package strlib

import "testing"

func TestAc(t *testing.T) {
	patterns := []string{"你好", "很好", "还好"}

	m := NewAcMatcher(patterns)

	r := m.Replace([]byte("今天的天气还好,心情很好 , 还好今天休息!!"), []byte("A"))
	t.Logf("%s", string(r))
}
