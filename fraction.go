package main

import (
	"fmt"
	"os"

	"github.com/dcaponi/fraction_fun/fracarg"
)

func main() {
	var a, b, r string
	var op rune

	if len(os.Args) != 4 {
		if len(os.Args) > 4 {
			fmt.Println("hint: you should escape the * symbol (e.g. 1/2 \\* 1/2)")
		}
		fmt.Println("input must follow format: a operand b")
		return
	}

	a = os.Args[1]
	op = rune(os.Args[2][0])
	b = os.Args[3]
	fmt.Println(os.Args)
	fa := fracarg.New(a)
	fb := fracarg.New(b)

	switch op {
	case '+':
		r = fa.Add(fb).ToString()
	case '-':
		r = fa.Sub(fb).ToString()
	case '*', 'x':
		r = fa.Mul(fb).ToString()
	case '/':
		r = fa.Div(fb).ToString()
	default:
		fmt.Println("illegal operator: operator must be one of + - x or /")
		return
	}
	fmt.Printf("result: %s\n", r)
}
