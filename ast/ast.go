package ast

import "github.com/go-interpreter/token"

// ast node의 기본형
type Node interface {
	TokenLiteral() string
}

// StatementNode
type Statement interface { // 선언문
	Node
	statementNode()
}

// ExpressionNode
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

// StatementNode 중 하나인 LetStatementNode
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// StatementNode 중 하나인 ReturnStatementNode
type ReturnStatement struct {
	Token token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { 
	return rs.Token.Literal
}

// ExpressionNode중 하나인 IdentifierNode
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
