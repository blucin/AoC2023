package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// i and j is the starting point S
	lines, i, j := ParseLines(ReadFile("input.txt"))
	t := Traverse(lines, i, j)
	solution1 := len(t) / 2
	fmt.Println("Solution 1:", solution1)
}

func Traverse(lines []string, i int, j int) []rune {
	traversed := make([]rune, 0)
	direction := ""
	for {
		if lines[i][j] == 'S' && len(traversed) > 0 {
			return traversed
		}
		// left
		if j != 0 && lines[i][j-1] != '.' && direction != "right" && lines[i][j] != 'L' && lines[i][j] != 'F' {
			direction = "left"
			for j >= 0 && lines[i][j-1] == '-' {
				traversed = append(traversed, '-')
				j--
			}
			if j != 0 && lines[i][j-1] != '-' && lines[i][j-1] != '|' {
				traversed = append(traversed, rune(lines[i][j-1]))
				j--
				continue
			}
		}
		// right
		if j != len(lines[i])-1 && lines[i][j+1] != '.' && direction != "left" && lines[i][j] != 'J' && lines[i][j] != '7' {
			direction = "right"
			for j < len(lines[i])-1 && lines[i][j+1] == '-' {
				traversed = append(traversed, '-')
				j++
			}
			if j != len(lines[i])-1 && lines[i][j+1] != '-' && lines[i][j+1] != '|' {
				traversed = append(traversed, rune(lines[i][j+1]))
				j++
				continue
			}
		}
		// up
		if i != 0 && lines[i-1][j] != '.' && direction != "down" && lines[i][j] != '7' && lines[i][j] != 'F' {
			direction = "up"
			for i >= 0 && lines[i-1][j] == '|' {
				traversed = append(traversed, '|')
				i--
			}
			if i != 0 && lines[i-1][j] != '|' && lines[i-1][j] != '-' {
				traversed = append(traversed, rune(lines[i-1][j]))
				i--
				continue
			}
		}
		// down
		if i != len(lines)-1 && lines[i+1][j] != '.' && direction != "up" && lines[i][j] != 'J' && lines[i][j] != 'L' {
			direction = "down"
			for i < len(lines)-1 && lines[i+1][j] == '|' {
				traversed = append(traversed, '|')
				i++
			}
			if i != len(lines)-1 && lines[i+1][j] != '|' && lines[i+1][j] != '-' {
				traversed = append(traversed, rune(lines[i+1][j]))
				i++
				continue
			}
		}
	}
}

func ChangeDirection(i int, j int, c rune) (int, int) {
	fmt.Println("ChangeDirection", i, j, c)
	switch c {
	case 'J':
		return i, j - 1
	case 'F':
		return i, j + 1
	case '7':
		return i + 1, j
	case 'L':
		return i - 1, j
	}
	return i, j
}

func ParseLines(lines []string) ([]string, int, int) {
	i, j := 0, 0
	for x, line := range lines {
		if i == 0 && j == 0 {
			for y, char := range line {
				if char == 'S' {
					i = x
					j = y
					break
				}
			}
		}
		lines = append(lines, line)
	}
	return lines, i, j
}

func ReadFile(fileName string) []string {
	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
