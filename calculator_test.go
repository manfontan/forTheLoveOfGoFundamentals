package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{inputs: []float64{2, 2}, want: 4},
		{inputs: []float64{1, 1, 1}, want: 3},
		{inputs: []float64{5, 0, 0, 0}, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		inputs []float64
		want   float64
	}

	testCases := []testCase{
		{inputs: []float64{2, 1}, want: 1},
		{inputs: []float64{6, -2}, want: 8},
		{inputs: []float64{12, -1, 1, 0}, want: 12},
		{inputs: []float64{12, 0.25, 0.25}, want: 11.5},
	}

	for _, tc := range testCases {
		got := calculator.Substract(tc.inputs...)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		inputs []float64
		want   float64
		name   string
	}

	testCases := []testCase{
		{inputs: []float64{2, 1}, want: 2, name: "Multiply to positive numbers"},
		{inputs: []float64{2, 6, -1}, want: -12, name: "Multiply a positive and a negative number"},
		{inputs: []float64{12, 0}, want: 0, name: "Multiply by 0"},
		{inputs: []float64{12, 0.5, 1, 1}, want: 6, name: "Multiply a positive number and a fractional number"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.inputs...)
		if tc.want != got {
			t.Errorf("%s : want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		inputs      []float64
		want        float64
		errExpected bool
		name        string
		err         string
	}

	testCases := []testCase{
		{inputs: []float64{1, 0, 1},
			want:        0,
			errExpected: true,
			name:        "division by zero",
			err:         "Divide(%f): unexpected error status %v"},
		{inputs: []float64{1, 1, 0},
			want:        0,
			errExpected: true,
			name:        "division by zero",
			err:         "Divide(%f): unexpected error status %v"},
		{inputs: []float64{8, 4, 2, 1},
			want:        1,
			errExpected: false,
			name:        "division",
			err:         "Divide(%f): unexpected error status %v",
		},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.inputs...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf(tc.err,
				tc.inputs, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("got %f want %f", got, tc.want)
		}
	}
}

func TestAddRand(t *testing.T) {
	rand.Float64()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	a := rand.Float64()
	want := math.Sqrt(a)
	got, err := calculator.Sqrt(a)
	errExpected := false

	errReceived := err != nil

	if errReceived != errExpected {
		t.Fatalf(" Sqrt(%f): Unexpected error status > %v", a, err)
	}

	if !errExpected && got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		expr        string
		want        float64
		errExpected bool
	}

	testCases := []testCase{
		{expr: "2 * 2", want: 4, errExpected: false},
		{expr: "1 + 1.5", want: 2.5, errExpected: false},
		{expr: "18  /  6", want: 3, errExpected: false},
		{expr: " 100 - 0.1 ", want: 99.9, errExpected: false},
		{expr: "2+2", want: 0, errExpected: true},
		{expr: "8 * 3 / 9", want: 0, errExpected: true},
		{expr: "x + 19", want: 0, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.expr)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Unexpected exception %v", err)
		}
		if !tc.errExpected && got != tc.want {
			t.Errorf("got %f want %f", got, tc.want)
		}
	}
}
