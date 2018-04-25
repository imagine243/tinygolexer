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
	var token *Token
	var err error
	for lexer.Index() < lexer.Count() {
		for _, m := range lexer.matchers {
			token, err = m.Match(lexer)
			if err != nil {
				log.Fatal(err)
			}
			if token != nil {
				if wordIndex > 0 {
					wordRaw := lexer.PeekChunk(lexer.Index()-wordIndex-token.length, lexer.Index()-token.length)
					wordToken := NewToken(nil, lexer.Index()-token.length, wordIndex, string(wordRaw), wordRaw)
					lexer.tokens = append(lexer.tokens, *wordToken)
					wordIndex = 0
				}
				lexer.tokens = append(lexer.tokens, *token)
				break
			}
		}
		if token == nil {
			wordIndex++
			fmt.Println("wordIndex : ", wordIndex)
			lexer.ConsumeMulti(1)
		}
		token, err = nil, nil
	}
	if wordIndex > 0 {
		wordRaw := lexer.PeekChunk(lexer.Index()-wordIndex, lexer.Index())
		wordToken := NewToken(nil, lexer.Index(), wordIndex, string(wordRaw), wordRaw)
		lexer.tokens = append(lexer.tokens, *wordToken)
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
	if begin >= lexer.Count() {
		var temp []rune
		return temp
	}
	if end >= lexer.Count() {
		return lexer.src[begin:]
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
