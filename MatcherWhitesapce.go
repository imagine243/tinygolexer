package tinygolexer

import "reflect"

type MatcherWhiteSpace struct {
	baseMatcher
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'

}

func (matcher *MatcherWhiteSpace) Match(lexer *Lexer) (*Token, error) {
	var index int
	for {
		c := lexer.Peek(index)
		if !isWhiteSpace(c) {
			break
		}
		index++
	}

	if index == 0 {
		return nil, nil
	}

	lexer.ConsumeMulti(index)
	return NewToken(matcher, lexer, " ", ""), nil
}

func (matcher *MatcherWhiteSpace) String() string {
	return reflect.TypeOf(matcher).Elem().Name()
}

func NewMatcherWhiteSpace() matcher {
	return &MatcherWhiteSpace{}
}
