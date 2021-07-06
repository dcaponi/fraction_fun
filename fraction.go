package main

import (
	"fmt"
	"os"

	"github.com/dcaponi/fraction_fun/fracarg"
)

func main() {
	var a, b string
	var y fracarg.FracArg
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

	fa := fracarg.New(a)
	fb := fracarg.New(b)

	switch op {
	case '+':
		y = fa.Add(fb)
	case '-':
		y = fa.Sub(fb)
	case '*', 'x':
		y = fa.Mul(fb)
	case '/':
		y = fa.Div(fb)
	default:
		fmt.Println("illegal operator: operator must be one of + - x or /")
		return
	}
	y.Red()
	fmt.Printf("result: %s\n", y.ToString())
}
