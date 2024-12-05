package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func SumBreakingRules(rules [][2]int, protocol []int, depth int) int {
	anyBreakingRules := false
	breakingRules := make([][2]int, 0)

	for i, num := range protocol {
		if i == 0 {
			continue
		}

		for j, rule := range rules {
			if rule[0] == num {
				foundIndex := slices.Index(protocol[0:i], rule[1])
				if foundIndex != -1 {
					fmt.Println(protocol, " breaks rule: ", rule, " at ", j)
					breakingRules = append(breakingRules, rule)
					anyBreakingRules = true
				}
			}
		}
	}

	fmt.Println("Pre swap: ", protocol)
	for _, rule := range breakingRules {
		i1 := slices.Index(protocol, rule[0])
		i2 := slices.Index(protocol, rule[1])
		protocol[i1], protocol[i2] = protocol[i2], protocol[i1]
	}
	fmt.Println("Post swap: ", protocol)

	if len(breakingRules) != 0 {
		return SumBreakingRules(rules, protocol, depth+1)
	}

	if len(breakingRules) == 0 && (anyBreakingRules || depth > 0) {
		middleIndex := len(protocol) / 2
		return protocol[middleIndex]
	}
	return 0
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	rules := make([][2]int, 0, 100)
	protocols := make([][]int, 0, 100)

	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, ",") {
			protocol := make([]int, 0, 5)
			for _, num := range strings.Split(line, ",") {
				numerical, _ := strconv.Atoi(num)
				protocol = append(protocol, numerical)
			}
			protocols = append(protocols, protocol)
		}
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			first, _ := strconv.Atoi(split[0])
			second, _ := strconv.Atoi(split[1])
			rules = append(rules, [2]int{first, second})
		}
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		sum := 0
		for _, protocol := range protocols {
			sum += SumBreakingRules(rules, protocol, 0)
		}

		return sum
	}

	sum := 0
	for _, protocol := range protocols {
		breakingRules := false
		for i, num := range protocol {
			if i == 0 {
				continue
			}

			for j, rule := range rules {
				if rule[0] == num {
					if slices.Contains(protocol[0:i], rule[1]) {
						fmt.Println(protocol, " breaks rule: ", rule, " at ", j)
						breakingRules = true
						break
					}
				}
			}

			if breakingRules {
				break
			}
		}
		if !breakingRules {
			middleIndex := len(protocol) / 2
			fmt.Println(middleIndex, protocol[middleIndex])
			sum += protocol[middleIndex]
		}
	}

	// solve part 1 here
	return sum
}
