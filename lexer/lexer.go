package lexer

import (
	"meitei_lon/token"
)

type lexer struct{
   input []rune 
   inputLen int
   ch    rune
   position int   
   readPosition int
}

func newToken(tt token.TokenType, ch rune) token.Token{
    return token.Token{ Type: tt, Literal: string(ch) } 
}

func NewLexer(input string) *lexer{
    var runeInput = []rune(input)
    l := &lexer{input : runeInput, inputLen : len(runeInput)}
    l.readChar()
    return l
}

func (l *lexer) skipWhiteSpace() {
	 for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		 l.readChar()
	 }
}

func(l *lexer) readChar(){
   if l.readPosition >= l.inputLen {
      l.ch = 0 
   } else {
       l.ch = rune(l.input[l.readPosition])
   }
   l.position = l.readPosition
   l.readPosition += 1
}

func isLetter(ch rune) bool{
    var last rune = 44013  
    return last >= ch && 'ꯀ' <= ch || ch == '_'
}

func isDigit(ch rune) bool{
    return '꯹' >= ch && '꯰' <= ch 
}

func(l *lexer) readIdentifier() string{
    position := l.position
    for isLetter(l.ch){
        l.readChar()
    } 
    return string(l.input[position:l.position])
}
func(l *lexer) peekChar() rune{
    if l.readPosition >= l.inputLen{
        return 0
    }
    return l.input[l.readPosition]
}

func(l *lexer) readNumber() (string, string){
    var numberType string = token.INT 
    position := l.position   
    for isDigit(l.ch){
        l.readChar()
    }

    if l.peekChar() == 'f'{  numberType = token.FLOAT } 
    number := string(l.input[position:l.position]) 
    l.readChar()
    return number, numberType
}

func(l *lexer) NextToken() token.Token{
    tok := token.Token{}     
    l.skipWhiteSpace()

    switch l.ch{
    case '!':
        if l.peekChar() == '='{
            ch := l.ch
            l.readChar()
            tok = token.Token{Type : token.NOTEQ, Literal : string(ch)+string(l.ch)}
        } else {
            tok = newToken(token.BANG, l.ch)
        }
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case ':':
        tok = newToken(token.COLON, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
// Operators
    case '=':
        if l.peekChar() == '='{
            ch := l.ch 
            l.readChar()
            tok = token.Token{Type: token.EQ, Literal : string(ch) + string(l.ch)}
        } else {
            tok = newToken(token.ASSIGN, l.ch)
        }
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '/':
        tok = newToken(token.SLASH, l.ch)
    case '*':
        tok = newToken(token.ASTERIK, l.ch)
    case '>':
        tok = newToken(token.GT, l.ch)
    case '<':
        tok = newToken(token.LT, l.ch)
    case 0:
        tok.Type = token.EOF
        tok.Literal = ""
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookUpIdentifier(tok.Literal)
            return tok
        } else if isDigit(l.ch){
            var numType string
            tok.Literal, numType = l.readNumber()             
            tok.Type  = token.LookUpDataType(numType) 
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        } 
    }
    l.readChar()
    return tok
}

