package ast

import "github.com/go-interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface { // 선언문
	Node
	statementNode()
}

type Expression interface { // 표현식
	Node
	expressionNode()
}

// statement 전체를 담는 그릇
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Statement Node 중 LET
type LetStatement struct {
	Token token.Token // token.LET
	Name *Identifier
	Value Expression
}
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Expression Node 중 IDENT. (expression이지만 값을 만들지 않음. 노드 타입 수를 줄여 단순화를 위한 trade-off)
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

