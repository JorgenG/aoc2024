package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"aoc-in-go/internal/utility"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// solve part 1 here
	lines := strings.Split(input, "\n")
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for _, line := range lines {
		before, after, _ := utility.SplitAndTrim(line, " ")
		leftNum, _ := strconv.Atoi(before)
		rightNum, _ := strconv.Atoi(after)
		left = append(left, leftNum)
		right = append(right, rightNum)
		sort.Ints(left)
		sort.Ints(right)
	}

	numberCounts := make(map[int]int, 100)
	for _, value := range right {
		count, exists := numberCounts[value]
		if !exists {
			numberCounts[value] = 1
		}
		numberCounts[value] = count + 1
	}

	if part2 {
		multiSum := 0
		for i := range left {
			count, found := numberCounts[left[i]]
			if !found {
				continue
			}
			multiSum += left[i] * count
		}

		return multiSum
	}

	totalDiff := 0
	for i := range left {
		totalDiff += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDiff
}
