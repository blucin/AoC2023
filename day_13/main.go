package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	patterns := ParseLine(ReadFile("input_test.txt"))
	solution1 := CheckReflections(patterns, false)
	fmt.Println("Solution 1: ", solution1)
}

func CheckReflections(patterns [][]string, checkSmudge bool) int {
	sum := 0
	for _, pattern := range patterns {
	Skip1:
		// horizontal
		for i := 0; i < len(pattern)-1; i++ {
			if pattern[i] == pattern[i+1] {
				above := i
				below := i + 1
				for above >= 0 && below < len(pattern) {
					if pattern[above] == pattern[below] {
						above--
						below++
					} else {
						continue Skip1
					}
				}
				sum += 100 * (i + 1)
			}
		}
	Skip2:
		// vertical
		for i := 0; i < len(pattern[0])-1; i++ {
			left := ""
			right := ""
			for j := range pattern {
				left += string(pattern[j][i])
				right += string(pattern[j][i+1])
			}
			if left == right {
				leftIdx := i
				rightIdx := i + 1
				for leftIdx >= 0 && rightIdx < len(pattern[0]) {
					leftStr := ""
					rightStr := ""
					for j := range pattern {
						leftStr += string(pattern[j][leftIdx])
						rightStr += string(pattern[j][rightIdx])
					}
					if leftStr == rightStr {
						leftIdx--
						rightIdx++
					} else {
						continue Skip2
					}
				}
				sum += i + 1
			}
		}
	}
	return sum
}

func ParseLine(lines []string) [][]string {
	patterns := make([][]string, 0)
	pattern := make([]string, 0)
	for i, line := range lines {
		if line != "" {
			pattern = append(pattern, line)
		}
		if line == "" || i == len(lines)-1 {
			patterns = append(patterns, pattern)
			pattern = make([]string, 0)
		}
	}
	return patterns
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
