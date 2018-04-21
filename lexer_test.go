package tinygolexer

import "testing"

func TestParser(t *testing.T) {
	l := NewLexer()
	l.AddMatcher(NewMatcherWhiteSpace())
	l.AddMatcher(NewMatcherKey("%%"))
	l.Lex("%% %%")

	if len(l.tokens) == 3 {
		t.Log(l.tokens)
		t.FailNow()
	}
}
