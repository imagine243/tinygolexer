package tinygolexer

type matcher interface {
	Match(lexer *Lexer) (*Token, error)
	String() string
	SetId(id int)
	Id() int
}

type baseMatcher struct {
	id int
}

func (m *baseMatcher) SetId(id int) {
	m.id = id
}

func (m *baseMatcher) Id() int {
	return m.id
}
