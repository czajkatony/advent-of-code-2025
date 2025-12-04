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
	ranges := strings.Split(cleanString, ",")
	answer := 0
	secondAnswer := 0

	// For each range
	for _, singleRange := range ranges {
		nums := strings.Split(string(singleRange), "-")
		idx, _ := strconv.Atoi(nums[0])
		finalNum, _ := strconv.Atoi(nums[1])

		for idx <= finalNum {
			if isInvalid(idx) {
				answer += idx
			}
			if isInvalid2(idx) {
				secondAnswer += idx
			}

			idx++
		}
	}

	fmt.Printf("Part 1 Total Sum: %d\n", answer)
	fmt.Printf("Part 2 Total Sum: %d\n", secondAnswer)
}

func isInvalid(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)

	if length%2 != 0 {
		return false
	}

	mid := length / 2
	return str[:mid] == str[mid:]
}

func isInvalid2(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)
	idx := 1
	// while indx <= length/2
	for idx <= length/2 {
		if shouldAdd(str, int(idx)) {
			return true
		}
		idx++
	}
	return false
}

func shouldAdd(str string, chunkSize int) bool {
	if len(str)%chunkSize != 0 {
		return false
	}
	prevString := str[:chunkSize]
	for len(str) > 0 {
		// If this chunk isn't the same as the last chunk, return false
		if prevString != str[:chunkSize] {
			return false
		}
		str = str[chunkSize:]
	}
	return true
}
