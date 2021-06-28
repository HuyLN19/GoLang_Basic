package main

import (
	"fmt"
)

func inputInt(iptr *int) {
	_, err := fmt.Scanln(iptr)
	if err != nil {
		fmt.Println("input invalid, please input again...")
		inputInt(iptr)
	}
}

func main() {
	var t int
	fmt.Println("Enter a Integer: ")
	inputInt(&t)
	fmt.Println("t =", t)
}
