package lexer

import (
	"fmt"
	"meitei_lon/token"
	"testing"
)

func TestLexer(t *testing.T){
    input := `
    ꯑꯣꯏꯍꯜꯂꯣ  ꯑ  : ꯃꯁꯤꯡ = .꯵꯰  
    ꯃꯊꯧ  ꯐꯉ꯭ꯀ () { 
        ꯑꯣꯏꯍꯜꯂꯣ  ꯂꯨ   : ꯃꯁꯤꯡ  = ꯵.꯰ 
        ꯑ  = ꯑ  + ꯂꯨ  
        ꯑꯣꯏꯔꯕ ꯑ  > ꯂꯨ  {
            ꯍꯟꯈꯣ ꯆꯨꯝꯃꯤ 
        }  ꯑꯣꯏꯗꯔꯕ {
            ꯍꯟꯈꯣ ꯂꯥꯜꯂꯤ  
        }
    }
    ꯵  < ꯷   
    ꯷  == ꯷  
    ꯵  != ꯷   
    !-/*
    `
    l := NewLexer(input)

    fmt.Println("CODE", input)

    test := []struct{
        expectedTokenType token.TokenType
        expectedLiteral string
    }{
        { token.NEXT_LINE, "\n" },
        { token.LET, "ꯑꯣꯏꯍꯜꯂꯣ" },
        { token.IDENT, "ꯑ" },
        { token.COLON, ":" },
        { token.NUMBERS, "ꯃꯁꯤꯡ"},
        { token.ASSIGN, "=" },
        { token.FLOAT,  "꯰.꯵꯰"},
        { token.NEXT_LINE, "\n" },
        { token.FUNCTION, "ꯃꯊꯧ" },  
        { token.IDENT,  "ꯐꯉ꯭ꯀ" },
        { token.LPAREN, "(" },
        { token.RPAREN, ")" },
        { token.LBRACE, "{" },
        { token.NEXT_LINE, "\n" },
        { token.LET, "ꯑꯣꯏꯍꯜꯂꯣ" },
        { token.IDENT, "ꯂꯨ" },
        { token.COLON, ":" },
        { token.NUMBERS, "ꯃꯁꯤꯡ"},
        { token.ASSIGN, "=" },
        { token.FLOAT,  "꯵.꯰"},
        { token.NEXT_LINE, "\n" },
        { token.IDENT, "ꯑ" },
        { token.ASSIGN, "=" },
        { token.IDENT, "ꯑ" },
        { token.PLUS, "+" },
        { token.IDENT, "ꯂꯨ" },
        { token.NEXT_LINE, "\n" },
        { token.IF, "ꯑꯣꯏꯔꯕ"},   
        { token.IDENT, "ꯑ" },
        { token.GT, ">" },
        { token.IDENT, "ꯂꯨ" },
        { token.LBRACE, "{" },
        { token.NEXT_LINE, "\n" },
        { token.RETURN, "ꯍꯟꯈꯣ" },
        { token.TRUE, "ꯆꯨꯝꯃꯤ" },
        { token.NEXT_LINE, "\n" },
        { token.RBRACE, "}" },
        { token.ELSE, "ꯑꯣꯏꯗꯔꯕ" },
        { token.LBRACE, "{" },
        { token.NEXT_LINE, "\n" },
        { token.RETURN, "ꯍꯟꯈꯣ" },
        { token.FALSE, "ꯂꯥꯜꯂꯤ" },
        { token.NEXT_LINE, "\n" },
        { token.RBRACE, "}" },
        { token.NEXT_LINE, "\n" },
        { token.RBRACE, "}" },
        { token.NEXT_LINE, "\n" },
        { token.INT, "꯵" },
        { token.LT, "<" },
        { token.INT, "꯷" },
        { token.NEXT_LINE, "\n" },
        { token.INT, "꯷" },
        { token.EQ, "==" },
        { token.INT, "꯷" },
        { token.NEXT_LINE, "\n" },
        { token.INT, "꯵" },
        { token.NOTEQ, "!=" },
        { token.INT, "꯷" },
        { token.NEXT_LINE, "\n" },
        { token.BANG, "!" },
        { token.MINUS, "-" },
        { token.SLASH, "/" },
        { token.ASTERIK, "*" },
    }
    
    
    for i, tt:= range test {
        tok := l.NextToken() 
        fmt.Printf("%+v\n", tok)
        if tok.Type != tt.expectedTokenType{
           t.Fatalf("tests[%d] expect-type=%q got=%q", i, tt.expectedTokenType, tok.Type) 
        } 

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] expect-literal=%q got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }    

}
