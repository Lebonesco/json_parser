package main

import (
	"encoding/json"

	"github.com/Lebonesco/json_parser/lexer"
	"io/ioutil"
	"os"
	"path/filepath"
)

func isJSONString(s string) bool {
	var str string
	return json.Unmarshal([]byte(s), &str) == nil
}

func main() {
	if len(os.Args) < 2 {
		panic("no valid file name or path provided for file!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err.Error())
	}

	l := lexer.NewLexer(data)
	for {
		token := l.NewToken()
		if string(token.Lit) == "" { // end
			return
		}
	}
}
