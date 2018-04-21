package tinygolexer

type Token struct {
	value interface{}
	raw   string
	pos   int
	m     matcher
}

func NewToken(m matcher, lexer *Lexer, v string, r string) *Token {
	self := Token{
		value: v,
		raw:   r,
		pos:   lexer.Count(),
		m:     m,
	}

	return &self
}
