package interpreter

import (
	"fmt"
	"log"
	"n0s4/goporth/lexer"
)

type Stack struct {
	items []int
}

func (s *Stack) push(val int) {
	s.items = append(s.items, val)
}
func (s *Stack) pop() int {
	len := len(s.items)
	if len == 0 {
		log.Fatalf("popped on empty stack")
	}
	val := s.items[len-1]
	s.items = s.items[:len-1]
	return val
}

func Interpret(tokens []lexer.Token) {
	stack := Stack{}
	for _, tok := range tokens {
		switch tok.Op {
		case lexer.PUSH:
			stack.push(tok.Val)

		case lexer.PLUS:
			b := stack.pop()
			a := stack.pop()
			stack.push(a + b)

		case lexer.MINUS:
			b := stack.pop()
			a := stack.pop()
			stack.push(a - b)

		case lexer.PRINT:
			v := stack.pop()
			fmt.Println(v)

		default:
			panic(fmt.Sprintf("unrecognised token while interpreting: %v", tok))
		}
	}
}
