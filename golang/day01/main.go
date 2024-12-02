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
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	left := make([]int, 1000)
	right := make([]int, 1000)

	for i, line := range lines {
		nums := strings.Fields(line)

		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])

		left[i] = l
		right[i] = r
	}

	sort.Ints(left)
	sort.Ints(right)

	fmt.Printf("Part one:\n%v\n", part1(left, right))
	fmt.Printf("Part two:\n%v\n", part2(left, right))

	fmt.Printf("Time: %v\n", time.Since(start))
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
