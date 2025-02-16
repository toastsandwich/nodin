package main

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	EOF  = "EOF" // end of file
	WSPC = "WHITE SPACE"
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
	Delimiter = map[string]struct{}{";": {}}

	Operator = map[string]struct{}{
		"+": {}, "-": {}, "*": {}, "/": {}, "%": {}, "++": {}, "--": {}, // arithmatic operators

		"==": {}, "!=": {}, ">=": {}, "<=": {}, ">": {}, "<": {}, // comparision operators

		"&&": {}, "||": {}, "!": {}, // logical operators

		"&": {}, "|": {}, "^": {}, "&^": {}, "<<": {}, ">>": {}, // bitwise operators

		"=": {}, "+=": {}, "-=": {}, "*=": {}, "/=": {}, "%=": {}, "|=": {}, "&=": {}, "^=": {}, "<<=": {}, "=>>": {}, "&^=": {}, // assignment operators

		"::": {}, "...": {}, ".": {}, ":": {}, // specials
	}

	Keyword = map[string]struct{}{
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

		"unint": {}, "u8": {}, "u16": {}, "u32": {}, "u64": {},
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
	if unicode.IsSpace(rune(l.Read())) {
		if l.Read() == '\n' {
			l.Line++
		}
		return &Token{
			Type:  "WHSPC",
			Value: "White Space",
			Line:  l.Line,
		}
	}

	switch l.Read() {
	case 0:
		return &Token{
			Type: EOF,
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
	if _, ok := Operator[op]; ok {
		nxt := string(l.ReadNext())
		if _, ok := Operator[op+nxt]; ok {
			op += nxt
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
		return &Token{
			Type:  INTL,
			Value: num,
			Line:  l.Line,
		}
	}
	if unicode.IsLetter(rune(l.Read())) {
		str := l.readString()
		if _, ok := Keyword[str]; ok {
			return &Token{
				Type:  KWRD,
				Value: str,
				Line:  l.Line,
			}
		}
		return &Token{
			Type:  STRL,
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
	for unicode.IsLetter(rune(l.ReadNext())) {
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
	return string(dgt)
}
