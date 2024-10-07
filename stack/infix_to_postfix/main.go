package main

import (
	"fmt"
	stack "infix/stack"
)

// Precedence function for operator precedence
func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/", "%":
		return 2
	case "^":
		return 3
	}
	return 0
}

// IsOperator checks if a string is an operator
func isOperator(c string) bool {
	return c == "+" || c == "-" || c == "*" || c == "/" || c == "%" || c == "^"
}

func main() {
	a := "a+b*(c^d-e)^(f+g*h)-i"
	s := stack.NewStack()
	r := ""

	for _, v := range a {
		sv := string(v)

		if isOperator(sv) {
			for s.Size() > 0 && precedence(sv) <= precedence(s.Peek()) {
				r = r + s.Pop()
			}
			s.Push(sv)
		} else if sv == "(" {
			s.Push(sv)
		} else if sv == ")" {
			for s.Peek() != "(" {
				r = r + s.Pop()
			}
			s.Pop()
		} else {
			r = r + sv
		}

	}

	for {
		if s.Size() != 0 {
			r = r + s.Pop()
		} else {
			break
		}
	}

	fmt.Println("input --> ", a)
	fmt.Println("output --> ", r)
}
