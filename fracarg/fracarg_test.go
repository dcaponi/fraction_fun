package fracarg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		Input    string
		Expected FracArg
	}{
		"It parses a mixed number": {
			Input: "3_1/2", Expected: FracArg{Num: 7, Den: 2},
		},
		"It parses a whole number": {
			Input: "3", Expected: FracArg{Num: 3, Den: 1},
		},
		"It parses a fraction": {
			Input: "1/2", Expected: FracArg{Num: 1, Den: 2},
		},
		"It parses a negative mixed number": {
			Input: "-3_1/2", Expected: FracArg{Num: -7, Den: 2},
		},
		"It parses a negative whole number": {
			Input: "-3", Expected: FracArg{Num: -3, Den: 1},
		},
		"It parses a negative fraction": {
			Input: "-1/2", Expected: FracArg{Num: -1, Den: 2},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, New(test.Input))
		})
	}
}

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		Input    []FracArg
		Expected FracArg
	}{
		"It adds FracArgs": {
			Input:    []FracArg{{Num: 7, Den: 2}, {Num: 2, Den: 3}},
			Expected: FracArg{Num: 25, Den: 6},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Input[0].Add(test.Input[1]))
		})
	}
}

func TestSub(t *testing.T) {
	tests := map[string]struct {
		Input    []FracArg
		Expected FracArg
	}{
		"It subtracts FracArgs": {
			Input:    []FracArg{{Num: 7, Den: 2}, {Num: 2, Den: 3}},
			Expected: FracArg{Num: 17, Den: 6},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Input[0].Sub(test.Input[1]))
		})
	}
}

func TestMul(t *testing.T) {
	tests := map[string]struct {
		Input    []FracArg
		Expected FracArg
	}{
		"It multiplies FracArgs": {
			Input:    []FracArg{{Num: 7, Den: 2}, {Num: 2, Den: 3}},
			Expected: FracArg{Num: 14, Den: 6},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Input[0].Mul(test.Input[1]))
		})
	}
}

func TestDiv(t *testing.T) {
	tests := map[string]struct {
		Input    []FracArg
		Expected FracArg
	}{
		"It divides FracArgs": {
			Input:    []FracArg{{Num: 7, Den: 2}, {Num: 2, Den: 3}},
			Expected: FracArg{Num: 21, Den: 4},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Input[0].Div(test.Input[1]))
		})
	}
}

func TestToString(t *testing.T) {
	tests := map[string]struct {
		Input    FracArg
		Expected string
	}{
		"It parses a mixed number": {
			Input: FracArg{Num: 7, Den: 2}, Expected: "3_1/2",
		},
		"It reduces and parses to number": {
			Input: FracArg{Num: 14, Den: 4}, Expected: "3_1/2",
		},
		"It parses a whole number": {
			Input: FracArg{Num: 3, Den: 1}, Expected: "3",
		},
		"It parses a fraction": {
			Input: FracArg{Num: 1, Den: 2}, Expected: "1/2",
		},
		"It reduces a fraction": {
			Input: FracArg{Num: 2, Den: 4}, Expected: "1/2",
		},
		"It parses a negative mixed number": {
			Input: FracArg{Num: -7, Den: 2}, Expected: "-3_1/2",
		},
		"It parses a negative whole number": {
			Input: FracArg{Num: -3, Den: 1}, Expected: "-3",
		},
		"It parses a negative fraction": {
			Input: FracArg{Num: -1, Den: 2}, Expected: "-1/2",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Input.ToString())
		})
	}
}
