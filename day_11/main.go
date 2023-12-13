package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := ExpandGalaxy(ReadFile("input_test.txt"), 1)
	for _, line := range lines {
		fmt.Println(string(line))
	}
	solution1 := GetDistanceSum(GetGalaxyCords(lines))
	fmt.Println("Solution 1:", solution1)

	lines = ExpandGalaxy(ReadFile("input.txt"), 1000000)
	solution2 := GetDistanceSum(GetGalaxyCords(lines))
	fmt.Println("Solution 2:", solution2-82)
}

// returns the sum of all possible manhattan distance between two points
// considers the combination of points with no repetition
func GetDistanceSum(coords [][]int) int {
	sum := 0
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			sum += Abs(coords[i][0]-coords[j][0]) + Abs(coords[i][1]-coords[j][1])
		}
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetGalaxyCords(lines [][]rune) [][]int {
	coords := make([][]int, 0)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				coords = append(coords, []int{i, j})
			}
		}
	}
	return coords
}

// expands rows and columns with no galaxy
// returns count of total galaxy and updated galaxy
func ExpandGalaxy(lines [][]rune, scale int) [][]rune {
	// expand rows
	emptyRows := make([]int, 0)
	for i, line := range lines {
		hasGalaxy := false
		for _, char := range line {
			if char != '.' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			emptyRows = append(emptyRows, i)
		}
	}
	lines = AddEmptyRow(lines, emptyRows, scale)

	// expand columns
	emptyColumns := make([]int, 0)
	cw := len(lines[0])
	for i := 0; i < cw; i++ {
		hasGalaxy := false
		for _, line := range lines {
			if line[i] != '.' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			emptyColumns = append(emptyColumns, i)
		}
	}
	lines = AddEmptyColumn(lines, emptyColumns, scale)
	return lines
}

func AddEmptyRow(lines [][]rune, rows []int, scale int) [][]rune {
	emptyRow := make([]rune, len(lines[0]))
	for i := range emptyRow {
		emptyRow[i] = '.'
	}
	offset := 0
	for _, row := range rows {
		row += offset
		for k := 0; k < scale; k++ {
			lines = append(lines, emptyRow)
			for j := len(lines) - 1; j > row; j-- {
				lines[j], lines[j-1] = lines[j-1], lines[j]
			}
		}
		offset += scale
	}
	return lines
}

func AddEmptyColumn(lines [][]rune, columns []int, scale int) [][]rune {
	offset := 0
	for _, column := range columns {
		column += offset
		for i := range lines {
			for k := 0; k < scale; k++ {
				lines[i] = append(lines[i], '.')
				for j := len(lines[i]) - 1; j > column; j-- {
					lines[i][j], lines[i][j-1] = lines[i][j-1], lines[i][j]
				}
			}
		}
		offset += scale
	}
	return lines
}

func ReadFile(fileName string) [][]rune {
	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	lines := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	return lines
}
