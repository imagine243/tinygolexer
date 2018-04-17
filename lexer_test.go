package tinygolexer

import "testing"

func TestParser(t *testing.T) {
	l := NewLexer()
	l.Lex("ttt")
}
