package tinygolexer

type Token struct {
	raw     string
	value   string
	matcher []rune
	pos     int
}
