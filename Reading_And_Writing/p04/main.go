package main

import (
	"fmt"
	"os"
)

func main() {
	inputFileReader, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("An error occurred on opening the input file: ", err)
		return
	}
	defer inputFileReader.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(inputFileReader, &v1, &v2, &v3)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("An error occurred on reading the input: ", err)
			return
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println("col1: ", col1)
	fmt.Println("col2: ", col2)
	fmt.Println("col3: ", col3)
}
