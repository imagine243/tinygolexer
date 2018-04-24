package tinygolexer

import (
	"fmt"
	"log"
)

// Lexer 词法分析器
type Lexer struct {
	src      []rune
	tokens   []Token
	matchers []matcher
	index    int
}

// Lex 词法分析
func (lexer *Lexer) Lex(s string) {
	lexer.src = []rune(s)
	lexer.index = 0
	fmt.Println(lexer.src)

	var wordIndex int
	for lexer.Index() < lexer.Count() {
		for _, m := range lexer.matchers {
			token, err := m.Match(lexer)
			if token != nil {
				if wordIndex > 0 {
					wordToken := NewToken(nil, lexer, string(lexer.PeekChunk(lexer.Index()-wordIndex, lexer.Index()-1)), "")
					lexer.tokens = append(lexer.tokens, *wordToken)
					wordIndex = 0
				}
				lexer.tokens = append(lexer.tokens, *token)
			}

			if err != nil {
				log.Fatal(err)
			}
		}
		wordIndex++
		lexer.ConsumeMulti(1)

	}
}

func (lexer *Lexer) AddMatcher(m matcher) {
	lexer.matchers = append(lexer.matchers, m)
}

func (lexer *Lexer) Index() int {

	return lexer.index
}

func (lexer *Lexer) Count() int {
	return len(lexer.src)
}

func (lexer *Lexer) Peek(offset int) rune {
	if lexer.index+offset >= len(lexer.src) {
		return 0
	}
	return lexer.src[lexer.index+offset]
}

func (lexer *Lexer) PeekChunk(begin int, end int) []rune {
	if begin >= lexer.Count() || end >= lexer.Count() {
		var temp []rune
		return temp
	}
	fmt.Printf("begin %d end %d\n", begin, end)
	return lexer.src[begin:end]
}

func (lexer *Lexer) ConsumeMulti(mul int) {
	lexer.index += mul
}

// NewLexer 构造词法分析器
func NewLexer() *Lexer {
	return &Lexer{}
}
