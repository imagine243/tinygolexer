package tinygolexer

type Token struct {
	value  interface{}
	raw    []rune
	pos    int
	length int
	m      matcher
}

func NewToken(m matcher, pos int, length int, v string, r []rune) *Token {
	self := Token{
		value:  v,
		raw:    r,
		pos:    pos - length,
		length: length,
		m:      m,
	}

	return &self
}
