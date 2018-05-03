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
	return NewToken(matcher, lexer.Index(), index, string(matcher.key), lexer.PeekChunk(lexer.Index()-index, lexer.Index())), nil
}

func (matcher *MatcherKey) String() string {
	return fmt.Sprintf("%s('%s')", reflect.TypeOf(matcher).Elem().Name(), string(matcher.key))
}

func NewMatcherKey(k string) matcher {
	return &MatcherKey{
		key: []rune(k),
	}
}

func NewMatcherKeyAndId(k string, id int) matcher {
	matcher := NewMatcherKey(k)
	matcher.SetId(id)
	return matcher
}
