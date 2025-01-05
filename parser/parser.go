package parser

import(
    "meitei_lon/lexer"
    "meitei_lon/token"
    "meitei_lon/ast"
    "fmt"
)

type Parser struct {
    l *lexer.Lexer
 
    curToken token.Token
    peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser{
    p := &Parser{l : l}

    p.nextToken()
    p.nextToken()

    return p
}

func(p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()

}

func(p *Parser) ParseProgram() *ast.Program{
    program := &ast.Program{}
    program.Statements = []ast.Statement{} 
    
    for p.curToken.Type != token.EOF{
        stmnt := p.parseStatement()
        if stmnt != nil{
            program.Statements = append(program.Statements, stmnt)
        }
        p.nextToken() 
    }

    return program
}

func(p *Parser) parseStatement() ast.Statement{
    switch p.curToken.Type{
    case token.LET:
        return p.parseLetStatement()
    default:
        return nil
    }
}

func(p *Parser) parseLetStatement() *ast.LetStatement {
    stmnt := &ast.LetStatement{ Token : p.curToken }

    if !p.expectPeek(token.IDENT){
       return nil 
    }

    stmnt.Name = &ast.Identifier{ Token: p.curToken, Value: p.curToken.Literal }

    if !p.expectPeek(token.ASSIGN){
       return nil 
    }
      // TODO 
      // condition for advancing the nextToken 
      // so that we can exit expression parsing
      // expression parsing

    fmt.Println("stmnt", stmnt) 
    return stmnt
}

func(p *Parser) curTokenIs(tt token.TokenType) bool{
   return p.curToken.Type == tt 
}

func(p *Parser) peekTokenIs(tt token.TokenType) bool{
   return p.peekToken.Type == tt 
}

func(p *Parser) expectPeek(tt token.TokenType) bool{
    if p.peekTokenIs(tt){
       p.nextToken()
       return true
    } else {
        return false
    }
}















