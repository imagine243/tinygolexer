package tinygolexer

import (
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) {
	l := NewLexer()
	l.AddMatcher(NewMatcherWhiteSpace())
	l.AddMatcher(NewMatcherKey("%%"))
	l.AddMatcher(NewMatcherFunc("{", "}"))
	l.AddMatcher(NewMatcherFunc("\"", "\""))
	l.AddMatcher(NewMatcherFunc("$", "$"))
	dat, err := ioutil.ReadFile("./testl.txt")
	if err != nil {
		panic(err)
	}
	testStr := string(dat)
	l.Lex(testStr)

	t.Log(l.tokens)
	t.FailNow()
}
