package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your City: ")
	city, _ := reader.ReadString('\n')
	fmt.Println("You live in ", city)
}
