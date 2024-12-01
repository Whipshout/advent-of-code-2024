package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	data, _ := os.ReadFile("../../inputs/01.txt")
	lines := strings.Split(string(data), "\n")

	left := []int{}
	right := []int{}

	for _, line := range lines {
		nums := strings.Fields(line)

		if len(nums) != 2 {
			continue
		}

		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	fmt.Printf("Part 1: %v\n", part1(left, right))
	fmt.Printf("Part 2: %v\n", part2(left, right))

	fmt.Printf("Execution time: %v\n", time.Since(start))
}

func part1(left, right []int) int {
	sumDistance := 0

	for i := 0; i < 1000; i++ {
		sumDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return sumDistance
}

func part2(left, right []int) int {
	frequency := make(map[int]int, 1000)

	for _, num := range right {
		frequency[num]++
	}

	similarScore := 0

	for _, num := range left {
		similarScore += frequency[num] * num
	}

	return similarScore
}
