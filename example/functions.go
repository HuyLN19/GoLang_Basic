package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c 
}

func vals() (int, int) {
	return 19, 12
}

func main() {
	res := plus(1, 2)
	fmt.Println("a + b =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("a + b + c = ", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_,c := vals()
	fmt.Println(c)
}

