package main

import "fmt"

type Student struct {
	name string
}

// Define Method
// func(t Type) methodName(params) returns {//body code}
// t Type => Receiver
// 1.Value Receiver
// 2.Pointer Receiver
func (s Student) getName() string {
	return s.name
}

// 1.Value Receiver
func (s Student) changeName() {
	fmt.Printf("p2 = %p\n", &s)
	s.name = "Zoro"
}

// 2.Pointer Receiver
func (s *Student) changeName2() {
	fmt.Printf("p2 = %p\n", s)
	s.name = "Zoro"
}

// non-struct
type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s, %s", s, str)
}

func main() {
	s := &Student{"Nguyen Huu Hung"}
	fmt.Println(s.getName())
	fmt.Printf("p1 = %p\n", s)

	s.changeName2()
	fmt.Println(s.name)
}
