package tinygolexer

type matcher interface {
	Match(lexer *Lexer) (*Token, error)
	String() string
}

type baseMatcher struct {
}
