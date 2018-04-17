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

// NewLexer 构造词法分析器
func NewLexer() *Lexer {
	return &Lexer{}
}
