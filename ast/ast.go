package main

import(
    "meitei_lon/token"
)

type Node interface(){
    TokenLiteral() string
}

type Statement interface{
    Node 
    statementNode()
}

type Expression interface{
    Node 
    expressionNode()
}

type Program struct {
    Statements []Statement
}

func(p *Program) TokenLiteral() string{
    if len(p.Statements) > 0{
       return p.Statements[0].TokenLiteral() 
    }
    return ""
}

/* ==================
    Identifier 
    ================
*/
type Identifier struct {
    Token token.Token
    Value string
}
func(i *Identifier) expressionNode(){}
func(i *Identifier) TokenLiteral() string{
    i.Token.Literal
}


/*
    ===================
    Let Statement 
    ==================
*/
type LetStatement struct {
    Token token.Token     
    Name *Identifier 
    Value Expression
}

func(l *LetStatement) statementNode(){}

func(l *LetStatement) TokenLiteral() string{
    l.Token.Literal
}






