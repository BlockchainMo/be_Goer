package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa int    = 123
	bb string = "s"
	cc bool   = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d,%q\n", a, s)
}

func InitVariableValue() {
	var a, b int = 3, 4
	var s = "BIOS"
	fmt.Printf("%d,%d,%s\n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c = 2, 5, true
	fmt.Println(a, b, c)
}

func variableShorter() {
	a, b, c := 2, 5, true
	fmt.Println(a, b, c)
}

func euler() {
	var c = 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a int = 3
	var b int = 4

	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}

func consts() {
	const a = 3
	const b = 4
	fmt.Println(math.Sqrt(a*a + b*b))
}

func enum() {
	const (
		c = iota
		_
		java
		python
		golang
	)
	fmt.Println(c, java, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	InitVariableValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)

	euler()
	consts()
	enum()

}
