package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinanry(input int) string {
	temp := ""
	for ; input > 0; input /= 2 {
		//temp += strconv.Itoa(input % 2)
		temp = temp + strconv.Itoa(input%2)
	}

	return temp
}

func converBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBinanry(15))
	fmt.Println(converBin(10))

	printFile("abc.txt")

	forever()
}
