package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file
	var numbers []float64
	file, err := os.Open("expenses.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read every line of the file, convert it to a float64 and add to the numbers slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(find2020(numbers))
}

func find2020(expenses []float64) float64 {
	var firstNum float64
	var secondNum float64
	var solution float64

	// outer for loop goes through the expenses slice one at a time
	for i := 0; i < len(expenses); i++ {
		firstNum = expenses[i]
		// inner loop adds every number in the slice to the firstNum until
		for i := 0; i < len(expenses); i++ {
			secondNum = expenses[i]
			if firstNum+secondNum == 2020 {
				// multiply the numbers together because that's the number Advent of Code wants
				solution = firstNum * secondNum
				break
			}
		}
	}
	return solution
}
