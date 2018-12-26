package ast

// base json Node
type Json interface {
	TokenLiteral()
}

type Array struct {
	Values []Json
}

type Object struct {
	Properties map[string]Json
}

type String struct {
	Value string
}

type Integer struct {
	Value string
}

// construct interface
func (node String) TokenLiteral()  {}
func (node Integer) TokenLiteral() {}
func (node Array) TokenLiteral()   {}
func (node Object) TokenLiteral()  {}
