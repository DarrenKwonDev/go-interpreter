package parser

import (
	"fmt"

	"github.com/go-interpreter/ast"
	"github.com/go-interpreter/lexer"
	"github.com/go-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // 현재 토큰
	peekToken token.Token // 다음 토큰
	errors []string
}

// 파서를 생성합니다.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken() 
	p.nextToken() // 현재, 다음 토큰을 얻기 위해서 nextToken 메서드 2회 실행
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// 현재 토큰을 기반으로 statement를 파싱합니다.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
		case token.LET: // LET 토큰을 만나면 LetStatement를 파싱합니다.
			return p.parseLetStatement()
		default:
			return nil
	}
}

// LET 토큰을 만나면 *ast.LetStatement 노드를 만듭니다.
// 우리는 let [IDENT] = [EXPRESSION]; 형태로 작성될 것을 기대하고 있습니다.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// 다음 토큰이 identifier가 와야 함.
	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	// let statement의 '이름'
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// IDENT 토큰 다음에는 값을 할당하는 할당 토큰(ASSIGN)이 올 것으로 기대함.
	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	// 세미콜론이 돌 때까지 일단은 건너 뛰자.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}