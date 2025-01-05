package token

type TokenType string

type Token struct {
    Type TokenType 
    Literal string   
}

var keywords = map[string]TokenType{

  "ꯑꯣꯏꯍꯜꯂꯣ"  : LET,  
  "ꯃꯊꯧ"     : FUNCTION,
  "ꯆꯨꯝꯃꯤ"    : TRUE,
  "ꯂꯥꯜꯂꯤ"    : FALSE,
  "ꯍꯟꯈꯣ"    : RETURN,
  "ꯑꯣꯏꯔꯕ"   : IF,
  "ꯑꯣꯏꯗꯔꯕ"  : ELSE,

  // DATA TYPE
  "ꯃꯁꯤꯡ"    : NUMBERS,
}

var dataType = map[string]TokenType{
    "INT"   : INT,
    "FLOAT" : FLOAT,
}

func LookUpIdentifier(identifier string) TokenType{
    if tt, exist := keywords[identifier]; !exist{
        return IDENT
    } else {
        return tt
    }
}

func LookUpDataType(t string) TokenType{
    if tt, exist := dataType[t]; !exist{
        return INVALID_DATA_TYPE 
    } else {
        return tt
    }
}

const (

    ILLEGAL = "ILLEGAL"
    INVALID_DATA_TYPE = "INVALID_DATA_TYPE"
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT = "INT" // 1343456
    FLOAT = "FLOAT" // 369.69
    
    // DATA TYPE
    NUMBERS = "NUMBERS"

    // Operators
    ASSIGN  = "="
    PLUS    = "+"
    MINUS   = "-"
    ASTERIK = "*"
    BANG    = "!"
    SLASH   = "/"

    LT      = "<"
    GT      = ">"
    EQ      = "=="
    NOTEQ   = "!="
    

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    COLON = ":"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
    IF          = "IF"
    ELSE        = "ELSE"
    RETURN      = "RETURN"

)
