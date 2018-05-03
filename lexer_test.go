package tinygolexer

import (
	"io/ioutil"
	"testing"
)

const (
	SPLIT = iota
)

func TestGenerate(t *testing.T) {
	l := NewLexer()
	l.AddMatcher(NewMatcherWhiteSpace())
	l.AddMatcher(NewMatcherKeyAndId("%%", SPLIT))
	l.AddMatcher(NewMatcherKey(":"))
	l.AddMatcher(NewMatcherKey("|"))
	l.AddMatcher(NewMatcherKey(";"))
	l.AddMatcher(NewMatcherFunc("{", "}"))
	l.AddMatcher(NewMatcherFunc("%{", "}%"))
	dat, err := ioutil.ReadFile("./testl.txt")
	if err != nil {
		panic(err)
	}

	testStr := string(dat)
	l.Lex(testStr)

	Generate("testl", l.tokens)
	t.FailNow()
}

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

func TestParserY(t *testing.T) {
	l := NewLexer()
	l.AddMatcher(NewMatcherWhiteSpace())
	l.AddMatcher(NewMatcherKey("%%"))
	l.AddMatcher(NewMatcherKey(":"))
	l.AddMatcher(NewMatcherKey("|"))
	l.AddMatcher(NewMatcherKey(";"))
	l.AddMatcher(NewMatcherFunc("{", "}"))
	dat, err := ioutil.ReadFile("./testy.txt")
	if err != nil {
		panic(err)
	}

	testStr := string(dat)
	l.Lex(testStr)

	t.Log(l.tokens)
	t.FailNow()
}
