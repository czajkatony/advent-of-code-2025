package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Get file content
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// convert bytes to string and split by newline
	cleanString := strings.TrimSpace(string(content))
	banks := strings.Split(cleanString, "\n")
	answer := 0
	answerPt2 := 0

	strDebug := "123456789"
	// 123
	fmt.Printf(strDebug[:3])
	fmt.Printf("\n")
	// 456789
	fmt.Printf(strDebug[3:])
	fmt.Printf("\n")

	// For each ranges
	for _, bank := range banks {
		answer += getJoltage(bank)
		answerPt2 += getJoltage2(bank)
	}

	fmt.Printf("Total Joltage: %d\n", answer)
	fmt.Printf("Part 2: Total Joltage: %d", answerPt2)
}

func getJoltage(bank string) int {
	// Start at idx 0 to len(bank) - 2 (Don't want to take the last value)
	startingTens, _ := strconv.Atoi(string(bank[0]))
	startingOnes, _ := strconv.Atoi(string(bank[1]))
	bestNum := startingTens*10 + startingOnes

	for i := 0; i <= len(bank)-2; i++ {
		tenVal, _ := strconv.Atoi(string(bank[i]))
		tens := tenVal * 10
		for j := i + 1; j <= len(bank)-1; j++ {
			oneVal, _ := strconv.Atoi(string(bank[j]))
			if tens+oneVal > bestNum {
				bestNum = tens + oneVal
			}
		}
	}
	return bestNum
}

func getJoltage2(bank string) int {
	return joltageHelper(bank, 12)
}

func joltageHelper(bank string, digitsLeft int) int {
	// Base case 0 return 0
	if digitsLeft == 0 {
		return 0
	}

	// Base case one will return one when we get largest digit with only one
	dlMinusOne := digitsLeft - 1
	digBuffer := len(bank) - dlMinusOne
	searchArea := bank[:digBuffer]
	fmt.Printf("Bank Length: %d "+"SearchArea: "+searchArea+" DigitsLeft: %d\n", len(bank), digitsLeft)
	largestDigit, ldIdx := getLargestDigit(searchArea)
	return largestDigit*getTens(dlMinusOne) + joltageHelper(bank[ldIdx:], dlMinusOne)
}

func getTens(numZeros int) int {
	if numZeros <= 0 {
		return 1
	}

	if numZeros <= 1 {
		return 10
	}

	num := 10
	for i := 1; i < numZeros; i++ {
		num *= 10
	}
	return num
}

func getLargestDigit(str string) (int, int) {
	if len(str) == 0 {
		fmt.Printf("String: " + str)
		panic("Shouldn't call getLargestDigit on empty string!!!")
	}

	largestDigit := 0
	idx := 0

	for i := 0; i < len(str); i++ {
		digit, _ := strconv.Atoi(string(str[i]))
		if digit > largestDigit {
			largestDigit = digit
			idx = i
		}
	}

	return largestDigit, idx + 1
}
