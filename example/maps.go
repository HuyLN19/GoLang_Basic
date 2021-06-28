package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 5
	m["k2"] = 11

	fmt.Println("map:", m)

	v := m["k1"]
	fmt.Println(v)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("map:",n)
}