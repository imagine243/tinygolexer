package tinygolexer

type Token struct {
	raw   string
	value string
	pos   int
	m     matcher
}

func NewToken(m matcher, lexer *Lexer, v string, r string) *Token {
	self := Token{
		value: v,
		raw:   r,
		m:     m,
	}

	return &self
}
