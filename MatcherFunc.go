package tinygolexer

import (
	"fmt"
	"reflect"
)

type MatcherFunc struct {
	baseMatcher
	leftSign  []rune
	rightSign []rune
	funcvalue string
}

func matcherSign(lexer *Lexer, beginIndex int, sign []rune) bool {
	var index int
	for _, c := range sign {
		if lexer.Peek(beginIndex+index) != c {
			return false
		}
		index++
	}
	return true
}

func (matcher *MatcherFunc) Match(lexer *Lexer) (*Token, error) {
	if (lexer.Count() - lexer.Index()) < (len(matcher.leftSign) + len(matcher.rightSign)) {
		return nil, nil
	}
	var index int
	if matcherSign(lexer, index, matcher.leftSign) {
		index += len(matcher.leftSign)
	} else {
		return nil, nil
	}

	for !matcherSign(lexer, index, matcher.rightSign) {
		index++
		if (lexer.Count() - index - lexer.Index()) < len(matcher.rightSign) {
			return nil, nil
		}
	}
	index += len(matcher.rightSign)
	fmt.Printf("lexer index : %d %d \n", lexer.Index(), index)
	v := string(lexer.PeekChunk(lexer.Index()+len(matcher.leftSign), lexer.Index()+index-len(matcher.rightSign)))
	r := lexer.PeekChunk(lexer.Index(), lexer.Index()+index)
	fmt.Printf("value %s raw %v \n", v, r)
	token := NewToken(matcher, lexer.Index(), index, v, r)
	lexer.ConsumeMulti(index)
	return token, nil
}

func (matcher *MatcherFunc) String() string {
	return fmt.Sprintf("%s ('%s %s %s')", reflect.TypeOf(matcher).Elem().Name(), string(matcher.leftSign), matcher.funcvalue, string(matcher.rightSign))
}

func NewMatcherFunc(leftSign string, rightSign string) matcher {
	return &MatcherFunc{
		leftSign:  []rune(leftSign),
		rightSign: []rune(rightSign),
	}
}
