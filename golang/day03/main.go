package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	data, _ := os.ReadFile("../../inputs/03.txt")
	text := string(data)

	fmt.Printf("Part one:\n%v\n", part1(text))
	fmt.Printf("Part two:\n%v\n", part2(text))

	fmt.Printf("Time: %v\n", time.Since(start))
}

func part1(text string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	sum := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}

func part2(text string) int {
	re := regexp.MustCompile(`(do\(\))|(don't\(\))|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	active := true
	sum := 0

	for _, match := range matches {
		switch {
		case match[1] != "":
			active = true
		case match[2] != "":
			active = false
		case active && match[3] != "" && match[4] != "":
			x, _ := strconv.Atoi(match[3])
			y, _ := strconv.Atoi(match[4])
			sum += x * y
		}
	}

	return sum
}
