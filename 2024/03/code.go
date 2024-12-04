package main

import (
	"strconv"
	"strings"

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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		fullInstructions := ""
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			fullInstructions += line
		}

		processedInstructions := fullInstructions

		for strings.Contains(processedInstructions, "don't()") {
			stripStart := strings.Index(processedInstructions, "don't()")
			stripEnd := strings.Index(processedInstructions[stripStart:], "do()")
			if stripEnd == -1 {
				processedInstructions = processedInstructions[:stripStart]
				break
			}
			processedInstructions = processedInstructions[:stripStart] + processedInstructions[stripEnd+stripStart:]
		}

		sum := 0
		_, remainder, found := strings.Cut(processedInstructions, "mul(")
		for ; found; _, remainder, found = strings.Cut(remainder, "mul(") {
			//		fmt.Println(remainder[0 : len(remainder)-1])
			firstNumber := 0
			separatorIndex := -1
			secondNumber := 0
			for i := 3; i > 0; i-- {
				number, err := strconv.Atoi(remainder[0:i])
				if err == nil {
					firstNumber = number
					separatorIndex = i
					break
				}
			}

			// 		fmt.Printf("%d %d\n", firstNumber, separatorIndex)

			if separatorIndex == -1 || remainder[separatorIndex:separatorIndex+1] != "," {
				continue
			}

			secondNumberIndex := separatorIndex + 1
			endIndex := -1

			for i := 3; i > 0; i-- {
				number, err := strconv.Atoi(remainder[secondNumberIndex : secondNumberIndex+i])
				if err == nil {
					secondNumber = number
					endIndex = i + secondNumberIndex
					break
				}
			}
			// 		fmt.Printf("%d %d\n", secondNumber, endIndex)

			if endIndex == -1 || remainder[endIndex:endIndex+1] != ")" {
				continue
			}

			sum += firstNumber * secondNumber
		}

		return sum
	}

	fullInstructions := ""
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		fullInstructions += line
	}

	sum := 0
	_, remainder, found := strings.Cut(fullInstructions, "mul(")
	for ; found; _, remainder, found = strings.Cut(remainder, "mul(") {
		//		fmt.Println(remainder[0 : len(remainder)-1])
		firstNumber := 0
		separatorIndex := -1
		secondNumber := 0
		for i := 3; i > 0; i-- {
			number, err := strconv.Atoi(remainder[0:i])
			if err == nil {
				firstNumber = number
				separatorIndex = i
				break
			}
		}

		// 		fmt.Printf("%d %d\n", firstNumber, separatorIndex)

		if separatorIndex == -1 || remainder[separatorIndex:separatorIndex+1] != "," {
			continue
		}

		secondNumberIndex := separatorIndex + 1
		endIndex := -1

		for i := 3; i > 0; i-- {
			number, err := strconv.Atoi(remainder[secondNumberIndex : secondNumberIndex+i])
			if err == nil {
				secondNumber = number
				endIndex = i + secondNumberIndex
				break
			}
		}
		// 		fmt.Printf("%d %d\n", secondNumber, endIndex)

		if endIndex == -1 || remainder[endIndex:endIndex+1] != ")" {
			continue
		}

		sum += firstNumber * secondNumber
	}

	// solve part 1 here
	return sum
}
