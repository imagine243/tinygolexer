package tinygolexer

type Token struct {
	value interface{}
	raw   string
	pos   int
	m     matcher
}

func NewToken(m matcher, lexer *Lexer, v string, r string) *Token {
	if r == "" {
		r = v
	}
	self := Token{
		value: v,
		raw:   r,
		pos:   lexer.Index(),
		m:     m,
	}

	return &self
}
