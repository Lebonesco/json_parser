package token

import "fmt"

type Token struct {
	Type
	Lit
	Pos
}

type Type string

type Lit []rune

type Pos struct {
	Offset int
	Line   int
	Column int
}

// Types
const (
	INVALID  = "INVALID"
	EOF      = "EOF"
	COMMA    = "COMMA"
	COLON    = "COLON"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"
	STRING   = "STRING"
	INTEGER  = "INTEGER"
)

func NewToken(typ Type, lit string) Token {
	return Token{Type: typ, Lit: []rune(lit)}
}

func (t *Token) String() string {
	return fmt.Sprintf("{token.%s, '%s'},", t.Type, string(t.Lit))
}
