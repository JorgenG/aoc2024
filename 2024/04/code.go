package main

import (
	"log"
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Nav struct {
	i int
	j int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// solve part 1 here
	if part2 {
		lines := strings.Split(input, "\n")
		cols := len(lines[0])
		rows := len(lines)

		matrix := make([][]int, rows)
		for i, line := range lines {
			matrix[i] = make([]int, cols)
			for j, char := range line {
				switch char {
				case []rune("M")[0]:
					matrix[i][j] = 1
				case []rune("A")[0]:
					matrix[i][j] = 2
				case []rune("S")[0]:
					matrix[i][j] = 3
				default:
					matrix[i][j] = 0
				}
			}
		}

		target1 := []int{1, 2, 3}
		target2 := []int{3, 2, 1}
		foundWords := 0
		for i, row := range matrix {
			for j := range row {
				if i > rows-2 || i < 1 {
					// Outside of bounds
					continue
				}

				if j > cols-2 || j < 1 {
					// Outside of bounds
					continue
				}

				criss := []int{
					matrix[i-1][j-1],
					matrix[i][j],
					matrix[i+1][j+1],
				}
				cross := []int{
					matrix[i+1][j-1],
					matrix[i][j],
					matrix[i-1][j+1],
				}

				if (slices.Equal(criss, target1) || slices.Equal(criss, target2)) && (slices.Equal(cross, target1) || slices.Equal(cross, target2)) {
					log.Printf("i%d j%d", i, j)
					foundWords++
				}
			}
		}

		return foundWords
	}

	lines := strings.Split(input, "\n")
	cols := len(lines[0])
	rows := len(lines)

	matrix := make([][]int, rows)
	for i, line := range lines {
		matrix[i] = make([]int, cols)
		for j, char := range line {
			switch char {
			case []rune("X")[0]:
				matrix[i][j] = 1
			case []rune("M")[0]:
				matrix[i][j] = 2
			case []rune("A")[0]:
				matrix[i][j] = 3
			case []rune("S")[0]:
				matrix[i][j] = 4
			default:
				matrix[i][j] = 0
			}
		}
	}

	navigationMatrix := []Nav{
		{i: 0, j: 1},   // Right
		{i: 1, j: 1},   // Down Right
		{i: 1, j: 0},   // Down
		{i: 1, j: -1},  // Down Left
		{i: 0, j: -1},  // Left
		{i: -1, j: -1}, // Left Up
		{i: -1, j: 0},  // Up
		{i: -1, j: 1},  // Up right
	}

	target := []int{1, 2, 3, 4}
	foundWords := 0
	for i, row := range matrix {
		for j := range row {
			for _, nav := range navigationMatrix {
				if nav.i*3+i > rows-1 || nav.i*3+i < 0 {
					// Outside of bounds
					continue
				}

				if nav.j*3+j > cols-1 || nav.j*3+j < 0 {
					// Outside of bounds
					continue
				}

				candidate := []int{
					matrix[i+nav.i*0][j+nav.j*0],
					matrix[i+nav.i*1][j+nav.j*1],
					matrix[i+nav.i*2][j+nav.j*2],
					matrix[i+nav.i*3][j+nav.j*3],
				}
				if slices.Equal(candidate, target) {
					foundWords++
				}
			}
		}
	}

	return foundWords
}
