package parser_test

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/parser"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
        let x = 5;
        let y = 10;
        let foobar = 838383;
    `

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokentLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%q", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
    return 5;
    return 10;
    return 696969420;
    `

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not a *ast.ReturnStatement. got=%q", returnStmt.TokenLiteral())
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("returnStmt.TokenLiteral() not 'return'. got=%q", returnStmt.TokenLiteral())
		}
	}
}

func checkParseErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
