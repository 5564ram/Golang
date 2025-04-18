package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReaderFromFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("An error occurred: ", err)
		return
	}
	defer inputReaderFromFile.Close()

	rearder := bufio.NewReader(inputReaderFromFile)
	for {
		input, err := rearder.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred: ", err)
			return
		}
		fmt.Printf("The input was: %s", input)

	}
}
