package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 3, 4
	fmt.Println(b, c)

	var d = true
	fmt.Println(d && false)

	var e int
	fmt.Println(e)

	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)

	apple := 2.0
	bread := 3.0
	price := apple + bread
	fmt.Printf("Price:     %.2f\n", price)

	home1 := "40 Pham Phu Tiet"
	home2 := "72 Nguyen Lai"
	home1, home2 = home2, home1
	fmt.Println("Home 1: ", home1)
	fmt.Println("Home 2: ", home2)

	f := true
	fmt.Println(f)
}
