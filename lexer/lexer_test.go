package lexer

import (
	"github.com/Lebonesco/json_parser/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `{
				    "glossary": {
				        "title": "example glossary",
						"GlossDiv": {
				            "title": "S",
							"GlossList": {
				                "GlossEntry": {
									"GlossTerm": "Standard Generalized Markup Language",
									"Abbrev": "ISO 8879:1986",
									"GlossDef": {
				                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
										"GlossSeeAlso": ["GML", "XML"]
				                    },
									"GlossSee": "markup"
				                }
				            },
				            "Nums": 5245243
				        }
				    }
				}`

	tests := []struct {
		typ token.Type
		lit string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "\"glossary\""},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "\"title\""},
		{token.COLON, ":"},
		{token.STRING, "\"example glossary\""},
		{token.COMMA, ","},
		{token.STRING, "\"GlossDiv\""},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "\"title\""},
		{token.COLON, ":"},
		{token.STRING, "\"S\""},
		{token.COMMA, ","},
		{token.STRING, "\"GlossList\""},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "\"GlossEntry\""},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "\"GlossTerm\""},
		{token.COLON, ":"},
		{token.STRING, "\"Standard Generalized Markup Language\""},
		{token.COMMA, ","},
		{token.STRING, "\"Abbrev\""},
		{token.COLON, ":"},
		{token.STRING, "\"ISO 8879:1986\""},
		{token.COMMA, ","},
		{token.STRING, "\"GlossDef\""},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "\"para\""},
		{token.COLON, ":"},
		{token.STRING, "\"A meta-markup language, used to create markup languages such as DocBook.\""},
		{token.COMMA, ","},
		{token.STRING, "\"GlossSeeAlso\""},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.STRING, "\"GML\""},
		{token.COMMA, ","},
		{token.STRING, "\"XML\""},
		{token.RBRACKET, "]"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "\"GlossSee\""},
		{token.COLON, ":"},
		{token.STRING, "\"markup\""},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "\"Nums\""},
		{token.COLON, ":"},
		{token.INTEGER, "5245243"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := NewLexer([]byte(input))

	for i, test := range tests {
		tok := l.NewToken()
		if test.typ != tok.Type {
			t.Fatalf("On test[%d], expected Type=%s, Got=%s", i, test.typ, tok.Type)
		}

		if test.lit != string(tok.Lit) {
			t.Fatalf("On test[%d], expected Literal=%s, Got=%s", i, test.lit, string(tok.Lit))
		}
	}
}
