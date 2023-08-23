package lexer

import (
	"fmt"
	"strconv"
	"unicode"
)

type Op uint8

const (
	PUSH Op = iota
	PLUS
	MINUS
	PRINT
)

func (op Op) String() string {
	switch op {
	case PUSH:
		return "push"
	case PLUS:
		return "plus"
	case MINUS:
		return "minus"
	case PRINT:
		return "print"
	default:
		panic("unrecognised operator when stringifying")
	}
}

type Token struct {
	Op  Op
	Val int // value if op == PUSH, otherwise ignore :/
}

func (tok Token) String() string {
	return fmt.Sprintf("{%v %d}", tok.Op, tok.Val)
}

func Lex(source string) ([]Token, error) {
	tokens := []Token{}
	var tokStart, tokEnd int
outer:
	for {
		// scan to next character
		for {
			if tokStart >= len(source) {
				break outer
			}
			if !unicode.IsSpace(rune(source[tokStart])) {
				break
			}
			tokStart++
		}

		tokEnd = tokStart
		for tokEnd < len(source) &&
			!unicode.IsSpace(rune(source[tokEnd])) {
			tokEnd++
		}

		tok, err := parseToken(source[tokStart:tokEnd])
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tok)

		if tokEnd == len(source)-1 {
			break
		} else {
			tokStart = tokEnd + 1
		}
	}
	return tokens, nil
}

func parseToken(token string) (Token, error) {

	if num, err := strconv.Atoi(token); err == nil {
		return Token{PUSH, num}, nil
	} else if err.(*strconv.NumError).Err == strconv.ErrRange {
		return Token{}, fmt.Errorf("int literal too big: %v", token)
	}

	switch token {
	case "+":
		return Token{Op: PLUS}, nil
	case "-":
		return Token{Op: MINUS}, nil
	case "print":
		return Token{Op: PRINT}, nil
	}

	return Token{}, fmt.Errorf("unexpected token: %s", token)
}
