package parser

import (
	"fmt"
	"meitei_lon/ast"
	"meitei_lon/lexer"
	"testing"
)

func TestLetStatement(t *testing.T){

    input :=`ꯑꯣꯏꯍꯜꯂꯣ ꯅꯝ = ꯵
    ꯑꯣꯏꯍꯜꯂꯣ ꯅꯝ_ = ꯷
    ꯑꯣꯏꯍꯜꯂꯣ ꯐꯨ_ꯕꯥꯔ = ꯶꯹꯱꯹꯳꯹꯱꯸꯳꯰`

    l := lexer.NewLexer(input) 
    p := NewParser(l)


    program := p.ParseProgram()
    if program == nil{
        t.Fatalf("ParseProgram returned nil")
    }
    if len(program.Statements) != 3{
       t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
    }
    
    tests := []struct {
        expectedIdentifier string
    }{
        { "ꯅꯝ" },
        { "ꯅꯝ_" }, 
        { "ꯐꯨ_ꯕꯥꯔ" },
    }

    for i, tt := range tests{
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier) {
            return   
        }
    } 
}

func testLetStatement(t *testing.T, s ast.Statement, name string ) bool{

    if s.TokenLiteral() != "ꯑꯣꯏꯍꯜꯂꯣ"{
        t.Errorf("s.TokenLiteral is not ' ꯑꯣꯏꯍꯜꯂꯣ ' got = %q", s.TokenLiteral()) 
        return false
    }

    letStmnt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s is not ast.Statement, got = %T", letStmnt)
        return false
    }

    if letStmnt.Name.Value != name{
        t.Errorf("letStmnt.Name.Value is not %q, got = %q", name, letStmnt.Name.Value)
        fmt.Println([]rune(letStmnt.Name.Value), []rune(name))
        return false
    }

    if letStmnt.Name.TokenLiteral() != name{
        t.Errorf("letStmnt.TokenLiteral is not %q, got = %q", name, letStmnt.TokenLiteral())
        return false
    }
    
    return true
}


