package main

import (
	"fmt"
	"os"

	"github.com/dcaponi/fraction_fun/fracarg"
)

func main() {
	var a, b, r string
	var op rune

	if len(os.Args) < 4 {
		fmt.Println("at least 3 arguments (a, operand, b) required")
		return
	}

	a = os.Args[1]
	op = rune(os.Args[2][0])
	b = os.Args[3]

	fa := fracarg.New(a)
	fb := fracarg.New(b)

	switch op {
	case '+':
		r = fa.Add(fb).ToString()
	case '-':
		r = fa.Sub(fb).ToString()
	case 'x':
		r = fa.Mul(fb).ToString()
	case '/':
		r = fa.Div(fb).ToString()
	default:
		fmt.Println("illegal operator: operator must be one of + - x or /")
		return
	}
	fmt.Printf("result: %s\n", r)
}
