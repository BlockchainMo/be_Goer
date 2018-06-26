//bool,string
//(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
//byte,rune
//float32,float64,complex64,complex128

package main

import (
	"fmt"
	"math"
)

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int

	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c, "1212")
}

func main() {
	consts()
}
