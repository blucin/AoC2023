package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := ReadFile("input.txt")
	solution1 := GetLoad(TiltNorth(lines))
	fmt.Println("Solution 1:", solution1)
}

func GetLoad(lines [][]rune) int {
	load := 0
	for i, line := range lines {
		for _, char := range line {
			if char == 'O' {
				load += len(lines) - i
			}
		}
	}
	return load
}

func TiltNorth(lines [][]rune) [][]rune {
	for i := len(lines) - 1; i > 0; i-- {
		for j, char := range lines[i] {
			if char == 'O' {
				if lines[i-1][j] == '.' {
					lines[i-1][j] = 'O'
					lines[i][j] = '.'
				}
				if lines[i-1][j] == 'O' && lines[i][j] == 'O' {
					k := i - 1
					for k > 0 && lines[k][j] != '.' {
						if lines[k][j] == '#' {
							break
						}
						k--
					}
					if lines[k][j] == '.' {
						lines[k][j] = 'O'
						lines[i][j] = '.'
					}
				}
			}
		}
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
