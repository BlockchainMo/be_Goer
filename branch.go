package main

import (
	"fmt"
	"io/ioutil"
)

int name

name int

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}

	fmt.Printf("%s", grade(10))
	fmt.Printf("%s", grade(101))

}
