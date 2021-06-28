package main

import "fmt"

// anonymous struct
type Student struct {
	id   int
	name string
}

// anonymous fields
type NoName struct {
	string
	int
}

type animalInfo struct {
	area     string
	category string
}

// struct long struct - nested struct
type Animal struct {
	id   int
	name string
	info animalInfo
}

func main() {
	// anonymous struct
	alan := Student{
		id:   123,
		name: "Alan",
	}
	fmt.Println(alan)

	zoro := Student{456, "Zoro"}
	fmt.Println(zoro)

	// Pointer tro toi struct
	robin := &Student{111, "Robin"}
	fmt.Println(*robin)
	// anonymous fields
	Huy := NoName{
		"Ngoc Huy",
		1912,
	}
	fmt.Println(Huy)

	// Nested struct
	pug := Animal{
		id:   12,
		name: "Ben",
		info: animalInfo{
			area:     "tang 2",
			category: "Dog",
		},
	}
	fmt.Println(pug)
}
