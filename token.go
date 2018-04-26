package tinygolexer

import (
	"fmt"
)

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

func (t *Token) String() string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("pos : %d length : %d value : %s raw : %v \n", t.pos, t.length, t.value, t.raw)
}
