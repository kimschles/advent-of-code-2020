package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() {
	// Open the data file for reading
	file, err := os.Open("smol-pw.txt")
	if err != nil {
		log.Fatal(err)
	}

	// create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// loops until the end of the file is reached and scanner.Scan returns false
	for scanner.Scan() {
		fmt.Printf("%T\n", scanner.Text())
	}

	// If there was an error scanning the file, report it and exit
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Close the file to free resources
	err = file.Close()

}

type searchTerms struct {
	lowNum       int
	highNum      int
	targetLetter string
	password     string
}

func stringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func stringSlicer() []searchTerms {

	samples := []string{"8-13 v: mtjkbnvvvhvvv", "8-14 k: xnjcftlkvhkmkr", "2-8 v: jvvvvvvjv"}

	var searchSlice []searchTerms

	for _, sample := range samples {
		pwString := strings.Split(sample, " ")
		targetLetter := strings.TrimRight(pwString[1], ":")
		numbers := strings.Split(pwString[0], "-")

		var tempStruct searchTerms
		tempStruct.lowNum = stringToInt(numbers[0])
		tempStruct.highNum = stringToInt(numbers[1])
		tempStruct.targetLetter = targetLetter
		tempStruct.password = pwString[2]

		searchSlice = append(searchSlice, tempStruct)

	}

	return searchSlice

}

func main() {
	pwMap := stringSlicer()
	for _, item := range pwMap {
		var targetLetter string = item.targetLetter
		var lowNum int = item.lowNum
		var highNum int = item.highNum
		var password string = item.password
		var letterMap = map[string]int{}
		var correctPassCount int

		for _, pw := range password {

			// count the number of times each letter is present in the string
			// store the count in a map

			for _, character := range strconv.QuoteRune(pw) {
				letterMap[string(character)]++

				//check if the number of the targetLetter is within the range
				if letterMap[targetLetter] >= lowNum && letterMap[targetLetter] <= highNum {
					correctPassCount++
				}

			}

		}
	}

}
