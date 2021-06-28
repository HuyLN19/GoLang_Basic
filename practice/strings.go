package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputInt(iptr *int) {
	_, err := fmt.Scanln(iptr)
	if err != nil {
		inputInt(iptr)
	}
}

func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func main() {
	fmt.Println(GetStrings("file.txt"))

	var t int
	fmt.Println("Enter a Integer: ")
	inputInt(&t)
	fmt.Println(t)
}
