package fracarg

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// FracArg is a fraction argument to the program in improper form
type FracArg struct {
	Num int
	Den int
}

// New converts string input into a FracArg by converting mixed numbers to
// improper fractions & finally FracArgs, or simply representing fraction
// inputs like 1/2 as FracArgs
func New(s string) FracArg {
	var w, n, d int
	var err error

	neg := s[0] == byte('-')
	if neg {
		s = s[1:]
	}

	// parse out the whole part if one exists
	parts := strings.Split(s, "_")
	if len(parts) > 1 {
		w, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("illegal argument given %s", s)
		}
	}

	parts = strings.Split(parts[len(parts)-1], "/")

	// no split happened, whole number only
	if len(parts) == 1 {
		w, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("illegal argument given %s", s)
		}
		if neg {
			w *= -1 // preserve sign
		}
		return FracArg{Num: w, Den: 1} // rep whole number as improper fraction
	}

	// numerator
	n, err = strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("illegal argument given %s", s)
	}

	// denominator
	d, err = strconv.Atoi(parts[1])
	if err != nil || d == 0 { // zero denominator protection
		log.Fatalf("illegal argument given %s", s)
	}

	// chuck the whole part onto the numerator
	n += w * d
	if neg {
		n *= -1 // preserve sign
	}
	return FracArg{Num: n, Den: d}
}

// Add adds FracArg b to FracArg a and return a FracArg
func (a FracArg) Add(b FracArg) FracArg {
	return FracArg{
		Num: (a.Num * b.Den) + (b.Num * a.Den),
		Den: (a.Den * b.Den),
	}
}

// Sub subtracts FracArg b from FracArg a and return a FracArg
func (a FracArg) Sub(b FracArg) FracArg {
	return FracArg{
		Num: (a.Num * b.Den) - (b.Num * a.Den),
		Den: (a.Den * b.Den),
	}
}

// Mul multiplies FracArg b and FracArg a and return a FracArg
func (a FracArg) Mul(b FracArg) FracArg {
	return FracArg{
		Num: (a.Num * b.Num),
		Den: (a.Den * b.Den),
	}
}

// Div divides FracArg a by FracArg b and return a FracArg
func (a FracArg) Div(b FracArg) FracArg {
	return FracArg{
		Num: (a.Num * b.Den),
		Den: (a.Den * b.Num),
	}
}

// Red reduces the FracArg to its simplest form
func (a *FracArg) Red() {
	g := gcd(int(math.Abs(float64(a.Num))), int(math.Abs(float64(a.Den))))
	a.Num /= g
	a.Den /= g
}

// ToString converts the FracArg to a string representation in the same style as
// the program inputs (e.g 3_1/3) in its reduced and/or mixed form (e.g. 7/2 becomes 3_1/2)
func (f FracArg) ToString() string {
	f.Red()
	if f.Num/f.Den == 0 {
		return fmt.Sprintf("%d/%d", f.Num, f.Den)
	}
	if f.Num%f.Den == 0 {
		return fmt.Sprintf("%d", f.Num/f.Den)
	}
	return fmt.Sprintf("%d_%d/%d", f.Num/f.Den, int(math.Abs(float64(f.Num%f.Den))), f.Den)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
