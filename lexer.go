package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType string

const (
	EOF  = "EOF" // end of file
	WSPC = "WHITE SPACE"
	FLTL = "FLOAT LITERAL"
	INTL = "INTEGER LITERAL"
	STRL = "STRING LITERAL"
	KWRD = "KEYWORD"
	IDNT = "IDENTIFIER"
	OPR  = "OPERATOR"
	OPAR = "OPEN PARENTHESIS"
	CPAR = "CLOSE PARENTHESIS"
	OBLK = "OPEN BLOCK"
	CBLK = "CLOSE BLOCK"
	DLMT = "DELIMITER"
	ILL  = "ILLEGAL"
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

func (t Token) String() string {
	return fmt.Sprintf("Token< Type: %s Line: %d Value: %s >", t.Type, t.Line, t.Value)
}

var (
	Delimiter = map[string]struct{}{";": {}, ":": {}, ",": {}}

	Operators = map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {}, "%": {}, "++": {}, "--": {}, // arithmatic operators

		"==": {}, "!=": {}, ">=": {}, "<=": {}, ">": {}, "<": {}, // comparision operators

		"&&": {}, "||": {}, "!": {}, // logical operators

		"&": {}, "|": {}, "^": {}, "&^": {}, "<<": {}, ">>": {}, // bitwise operators

		"=": {}, "+=": {}, "-=": {}, "*=": {}, "/=": {}, "%=": {}, "|=": {}, "&=": {}, "^=": {}, "<<=": {}, "=>>": {}, "&^=": {}, // assignment operators

		"::": {}, "...": {}, ".": {}, ":": {}, // specials
	}

	Keywords = map[string]struct{}{
		"pkg":       {},
		"pub":       {},
		"type":      {},
		"struct":    {},
		"interface": {},
		"var":       {},
		"func":      {},
		"goto":      {},
		"import":    {},

		"if": {}, "else": {}, "elif": {},

		"for": {},

		"int": {}, "i8": {}, "i16": {}, "i32": {}, "i64": {},

		"uint": {}, "u8": {}, "u16": {}, "u32": {}, "u64": {},

		"byte": {},

		"rune": {},

		"f32": {}, "f64": {},

		"str": {},

		"bool": {},
	}
)

/*
pkg main;
^^
func main() {
	var: i32 a = 100;
	var: bool b = a >= 10;
}
EOF
*/

type Lexer struct {
	Input string // code

	Current int
	Next    int
	Line    int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{Input: input, Line: 1}
	if len(input) > 1 {
		l.Next = 1
	}
	return l
}

func (l *Lexer) Read() byte {
	if l.Current >= len(l.Input) {
		return 0
	}

	return l.Input[l.Current]
}

func (l *Lexer) ReadNext() byte {
	if l.Next >= len(l.Input) {
		return 0
	}
	return l.Input[l.Next]
}

func (l *Lexer) Advance() {
	l.Current += 1
	l.Next = l.Current + 1
}

func (l *Lexer) ReadToken() *Token {
	defer l.Advance()
	for unicode.IsSpace(rune(l.Read())) {
		if l.Read() == '\n' {
			l.Line++
		}
		l.Advance()
	}

	switch l.Read() {
	case 0:
		return &Token{
			Type: EOF,
		}
	case '"':
		l.Advance()
		return &Token{
			Type:  STRL,
			Value: l.readStringLiteral(),
			Line:  l.Line,
		}
	case ';':
		return &Token{
			Type:  DLMT,
			Value: ";",
			Line:  l.Line,
		}
	case '{':

		return &Token{
			Type:  OBLK,
			Value: "{",
			Line:  l.Line,
		}
	case '}':
		return &Token{
			Type:  CBLK,
			Value: "}",
			Line:  l.Line,
		}
	case '(':
		return &Token{
			Type:  OPAR,
			Value: "(",
			Line:  l.Line,
		}
	case ')':
		return &Token{
			Type:  CPAR,
			Value: ")",
			Line:  l.Line,
		}
	}

	op := string(l.Read())
	if _, ok := Operators[op]; ok {
		nxt := string(l.ReadNext())
		if _, ok := Operators[op+nxt]; ok {
			op += nxt
			if l.Next+1 < len(l.Input) {
				nxtt := string(l.Input[l.Next+1])
				if _, ok := Operators[op+nxtt]; ok {
					op += nxtt
					l.Advance()
				}
			}
			l.Advance()
		}
		return &Token{
			Type:  OPR,
			Value: op,
			Line:  l.Line,
		}
	}

	if unicode.IsDigit(rune(l.Read())) {
		num := l.readDigit()
		if strings.Contains(num, ".") {
			return &Token{
				Type:  FLTL,
				Value: num,
				Line:  l.Line,
			}
		}
		return &Token{
			Type:  INTL,
			Value: num,
			Line:  l.Line,
		}
	}
	if unicode.IsLetter(rune(l.Read())) {
		str := l.readString()
		if _, ok := Keywords[str]; ok {
			return &Token{
				Type:  KWRD,
				Value: str,
				Line:  l.Line,
			}
		}
		return &Token{
			Type:  IDNT,
			Value: str,
			Line:  l.Line,
		}
	}

	return &Token{
		Type:  ILL,
		Value: string(l.Read()),
		Line:  l.Line,
	}
}

func (l *Lexer) readString() string {
	str := []byte{l.Read()}
	for unicode.IsLetter(rune(l.ReadNext())) || unicode.IsDigit(rune(l.ReadNext())) {
		l.Advance()
		str = append(str, l.Read())
	}
	return string(str)
}

func (l *Lexer) readDigit() string {
	dgt := []byte{l.Read()}
	for unicode.IsDigit(rune(l.ReadNext())) {
		l.Advance()
		dgt = append(dgt, l.Read())
	}
	if l.ReadNext() == '.' {
		l.Advance()
		dgt = append(dgt, '.')
		for unicode.IsDigit(rune(l.ReadNext())) {
			l.Advance()
			dgt = append(dgt, l.Read())
		}
	}
	return string(dgt)
}

func (l *Lexer) readStringLiteral() string {
	str := []byte{}
	for l.Read() != '"' && l.Read() != 0 {
		str = append(str, l.Read())
		l.Advance()
	}
	l.Advance()
	return string(str)
}
