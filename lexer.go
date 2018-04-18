package tinygolexer

import "fmt"

// Lexer 词法分析器
type Lexer struct {
	src    []rune
	tokens []Token
	index  int
}

// Lex 词法分析
func (lexer *Lexer) Lex(s string) {
	lexer.src = []rune(s)
	lexer.index = 0
	fmt.Println(lexer.src)
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

func (lexer *Lexer) ConsumeMulti(mul int) {
	lexer.index += mul
}

// NewLexer 构造词法分析器
func NewLexer() *Lexer {
	return &Lexer{}
}
