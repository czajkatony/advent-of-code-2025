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
	lines := strings.Split(string(content), "\n")

	// set starting value to 50
	lockValue := 50
	zeroHit := 0
	// Main execution loop
	for _, line := range lines {
		lockValue = performRotation(lockValue, line)
		if lockValue == 0 {
			zeroHit++
		}
	}
	fmt.Printf("# times zero was hit: %d", zeroHit)
}

func performRotation(startingValue int, line string) int {
	if line == "" || line == " " {
		return startingValue
	}

	direction := string(line[0])
	num, _ := strconv.Atoi(line[1:])

	return mathOp(direction, num, startingValue)
}

func mathOp(direction string, numRotations int, startingValue int) int {
	if direction != "L" && direction != "R" {
		panic("Direction should be L or R!!!!")
	} else if direction == "L" {
		sum := numRotations + startingValue
		for sum > 99 {
			sum = sum - 100
		}
		return sum
	} else {
		value := startingValue - numRotations
		for value < 0 {
			value = value + 100
		}
		return value
	}
}
