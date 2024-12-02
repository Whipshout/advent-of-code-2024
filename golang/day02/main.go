package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	data, _ := os.ReadFile("../../inputs/02.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	matrix := make([][]int, 1000)

	for i, line := range lines {
		rawNumbers := strings.Fields(line)

		numbers := make([]int, len(rawNumbers))

		for i, n := range rawNumbers {
			x, _ := strconv.Atoi(n)
			numbers[i] = x
		}

		matrix[i] = numbers
	}

	fmt.Printf("Part one:\n%v\n", part1(matrix))
	fmt.Printf("Part two:\n%v\n", part2(matrix))

	fmt.Printf("Time: %v\n", time.Since(start))
}

func part1(matrix [][]int) int {
	count := 0

	for _, numbers := range matrix {
		if validReport(numbers) {
			count++
		}
	}

	return count
}

func part2(matrix [][]int) int {
	count := 0

	for _, numbers := range matrix {
		if validReportDampened(numbers) {
			count++
		}
	}

	return count
}

// Much better than my Rust version
func validReport(numbers []int) bool {
	numbersLen := len(numbers)
	delta := 0

	for i := 0; i < numbersLen-1; i++ {
		diff := numbers[i] - numbers[i+1]

		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff > 0 {
			delta++
		} else if diff < 0 {
			delta--
		}
	}

	return delta == numbersLen-1 || delta == -(numbersLen-1)
}

// Could be improved, not doing it :D
func validReportDampened(numbers []int) bool {
	numbersLen := len(numbers)

	for i := 0; i < numbersLen; i++ {
		newNumbers := make([]int, 0, numbersLen-1)
		newNumbers = append(newNumbers, numbers[:i]...)
		newNumbers = append(newNumbers, numbers[i+1:]...)

		if validReport(newNumbers) {
			return true
		}
	}

	return false
}
