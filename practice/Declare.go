package main

import (
	"fmt"
	"reflect"
)

// Data types in GoLang
// int, uint, int8, uint8, bool, string, float32, float64, complex64, complex128, ...
// struct
// Array ([]int, [3]string, []struct{Name string},...)
// Maps (map[string]int)

func main() {

	var a string = "Huy Hoc Gioi"
	b := "Huy Học Giỏi"
	var c int = 3
	var d float64 = 4.7
	var e bool = true
	x := make(map[int]int)
	y := [5]int{1, 2, 3, 4, 5}
	z := make([]string, 4)

	x[1] = 3
	x[2] = 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(c, ", ", d, ", ", e)
	fmt.Println(x, ", ", reflect.TypeOf(x))
	fmt.Println(string(c), string(c), string(c))

}
