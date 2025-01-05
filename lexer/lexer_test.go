package lexer

import (
	"fmt"
	"meitei_lon/token"
	"testing"
)

func TestLexer(t *testing.T){
    input := `
    ꯑꯣꯏꯍꯜꯂꯣ  ꯑ  : ꯃꯁꯤꯡ = ꯵f
    ꯃꯊꯧ  ꯐꯉ꯭ꯀ () { 
        ꯑꯣꯏꯍꯜꯂꯣ  ꯂꯨ   : ꯃꯁꯤꯡ = ꯵f
        ꯑ  = ꯑ  + ꯂꯨ   
        ꯑꯣꯏꯔꯕ ꯑ  > ꯂꯨ  {
            ꯍꯟꯈꯣ ꯆꯨꯝꯃꯤ
        }  ꯑꯣꯏꯗꯔꯕ {
            ꯍꯟꯈꯣ ꯂꯥꯜꯂꯤ
        }
    }
    ꯵ < ꯷   
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
        { token.LET, "ꯑꯣꯏꯍꯜꯂꯣ" },
        { token.IDENT, "ꯑ" },
        { token.COLON, ":" },
        { token.NUMBERS, "ꯃꯁꯤꯡ"},
        { token.ASSIGN, "=" },
        { token.INT,  "꯵"},
        { token.FUNCTION, "ꯃꯊꯧ" },  
        { token.IDENT,  "ꯐꯉ꯭ꯀ" },
        { token.LPAREN, "(" },
        { token.RPAREN, ")" },
        { token.LBRACE, "{" },
        { token.LET, "ꯑꯣꯏꯍꯜꯂꯣ" },
        { token.IDENT, "ꯂꯨ" },
        { token.COLON, ":" },
        { token.NUMBERS, "ꯃꯁꯤꯡ"},
        { token.ASSIGN, "=" },
        { token.INT,  "꯵"},
        { token.IDENT, "ꯑ" },
        { token.ASSIGN, "=" },
        { token.IDENT, "ꯑ" },
        { token.PLUS, "+" },
        { token.IDENT, "ꯂꯨ" },
        { token.IF, "ꯑꯣꯏꯔꯕ"},   
        { token.IDENT, "ꯑ" },
        { token.GT, ">" },
        { token.IDENT, "ꯂꯨ" },
        { token.LBRACE, "{" },
        { token.RETURN, "ꯍꯟꯈꯣ" },
        { token.TRUE, "ꯆꯨꯝꯃꯤ" },
        { token.RBRACE, "}" },
        { token.ELSE, "ꯑꯣꯏꯗꯔꯕ" },
        { token.LBRACE, "{" },
        { token.RETURN, "ꯍꯟꯈꯣ" },
        { token.FALSE, "ꯂꯥꯜꯂꯤ" },
        { token.RBRACE, "}" },
        { token.RBRACE, "}" },
        { token.INT, "꯵" },
        { token.LT, "<" },
        { token.INT, "꯷" },
        { token.INT, "꯷" },
        { token.EQ, "==" },
        { token.INT, "꯷" },
        { token.INT, "꯵" },
        { token.NOTEQ, "!=" },
        { token.INT, "꯷" },
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
