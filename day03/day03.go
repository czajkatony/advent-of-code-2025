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

	// For each ranges
	for _, bank := range banks {
		fmt.Printf(bank)
		answer += getJoltage(bank)
		fmt.Printf(" totalAnswer: %d\n", answer)
	}

	fmt.Printf("Total Joltage: %d", answer)
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

	fmt.Printf(": bestNum --> %d", bestNum)
	return bestNum
}
