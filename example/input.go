package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Enter is : ", input)
}
