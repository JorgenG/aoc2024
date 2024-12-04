package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func isSafe(level []string) bool {
	current, _ := strconv.Atoi(level[1])
	prev, _ := strconv.Atoi(level[0])
	diff := current - prev

	minDiff := diff
	maxDiff := diff

	for i := range level {
		if i == 0 {
			continue
		}
		current, _ := strconv.Atoi(level[i])
		prev, _ := strconv.Atoi(level[i-1])
		diff := current - prev
		if minDiff > diff {
			minDiff = diff
		}
		if maxDiff < diff {
			maxDiff = diff
		}
	}

	if maxDiff >= 1 && maxDiff <= 3 && minDiff >= 1 {
		return true
	} else if minDiff <= -1 && minDiff >= -3 && maxDiff <= -1 {
		return true
	}
	return false
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	lines := strings.Split(input, "\n")

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		safeCounter := 0
		for _, line := range lines {
			level := strings.Split(line, " ")

			anySafe := false
			for i := range level {
				if isSafe(slices.Concat(level[:i], level[i+1:])) {
					anySafe = true
				}
			}
			if anySafe {
				safeCounter++
			}
		}

		return safeCounter
	}

	safeCounter := 0
	for _, line := range lines {
		level := strings.Split(line, " ")
		if isSafe(level) {
			safeCounter++
		}
	}

	// solve part 1 here
	return safeCounter
}
