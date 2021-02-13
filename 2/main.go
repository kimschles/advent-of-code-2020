package main

import "fmt"

// parse input

func main() {
	var targetLetter string = "j"
	var lowNum int = 3
	var highNum int = 4
	var password string = "hjujj"
	var letterMap = map[string]int{}
	var correctPassCount int

	// count the number of times each letter is present in the string
	// store the count in a map
	for _, character := range password {
		letterMap[string(character)]++
	}

	//check if the number of the targetLetter is within the range
	if letterMap[targetLetter] >= lowNum && letterMap[targetLetter] <= highNum {
		correctPassCount++
	}
	fmt.Println(correctPassCount)
}
