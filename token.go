package tinygolexer

type Token struct {
	raw     string
	value   string
	code    string
	matcher []rune
	pos     int
}
