package tinygolexer

import (
	"testing"
)

func TestParser(t *testing.T) {
	l := NewLexer()
	l.AddMatcher(NewMatcherWhiteSpace())
	l.AddMatcher(NewMatcherKey("%%"))
	l.AddMatcher(NewMatcherFunc("{", "}"))
	l.AddMatcher(NewMatcherFunc("\"", "\""))
	l.Lex("%% \"server\"        {return SERVER;} %%")

	if len(l.tokens) == 7 {
		t.Log(l.tokens)
		t.FailNow()
	}
}
