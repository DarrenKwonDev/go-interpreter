package parser

import (
	"github.com/go-interpreter/ast"
	"github.com/go-interpreter/lexer"
	"github.com/go-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // 현재 토큰
	peekToken token.Token // 다음 토큰
}

// 파서를 생성합니다.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken() 
	p.nextToken() // 현재, 다음 토큰을 얻기 위해서 nextToken 메서드 2회 실행
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}