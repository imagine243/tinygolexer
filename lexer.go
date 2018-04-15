package tinygolexer

type Lexer struct {
	src    []rune
	tokens *[]Token
	index  int
}

func (self *Lexer) Lex(s string) {
	self.src = []rune(s)
	self.tokens = new([]Token)
}
func NewLexer() *Lexer {

	return &Lexer{
		index: 0,
	}
}
