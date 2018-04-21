package tinygolexer

import (
	"fmt"
	"reflect"
)

type MatcherKey struct {
	baseMatcher
	key []rune
}

func (matcher *MatcherKey) Match(lexer *Lexer) (*Token, error) {
	if (lexer.Count() - lexer.Index()) < len(matcher.key) {
		return nil, nil
	}
	var index int
	for _, c := range matcher.key {
		if lexer.Peek(index) != c {
			return nil, nil
		}
		index++
	}

	lexer.ConsumeMulti(index)
	return NewToken(matcher, lexer, string(matcher.key), ""), nil
}

func (matcher *MatcherKey) String() string {
	return fmt.Sprintf("%s('%s')", reflect.TypeOf(matcher).Elem().Name(), string(matcher.key))
}

func NewMatcherKey(k string) matcher {
	return &MatcherKey{
		key: []rune(k),
	}
}
