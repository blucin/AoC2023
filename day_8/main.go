package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	lines := ReadFile("input.txt")
	traversal := lines[0]
	lines = append(lines[:0], lines[2:]...)

	solution1Keys := make([]string, 0)
	solution2Keys := make([]string, 0)

	mappings := make(map[string][2]string, 0)
	for i, line := range lines {
		mappings[line[0:3]] = [2]string{line[7:10], line[12:15]}
		if line[0:3] == "AAA" && len(solution1Keys) == 0 {
			solution1Keys = append(solution1Keys, lines[i][0:3])
		}
		if line[2] == 'A' {
			solution2Keys = append(solution2Keys, lines[i][0:3])
		}
	}
	fmt.Println(solution1Keys)
	fmt.Println(solution2Keys)

	solution1 := MaxTraversalCnt(`ZZZ`, mappings, solution1Keys, traversal)
	solution2 := MaxTraversalCnt(`..Z`, mappings, solution2Keys, traversal)

	fmt.Println("Solution 1:", solution1)
	fmt.Println("Solution 2:", solution2)
}

// we have a map of all the points and their connections either left or right
// returns max cnt of traversals it took to find a pattern in a map
// given indexs of initial arrays of key and a traversal string
func MaxTraversalCnt(pattern string, mappings map[string][2]string, keys []string, traversal string) int {
	counts := make([]int, 0)
	for _, key := range keys {
		start := key
		localCnt := 0
		for i := 0; ; i++ {
			t := traversal[i%len(traversal)]
			match, _ := regexp.MatchString(pattern, start)
			if match {
				fmt.Println("Found pattern", pattern, "at", start)
				break
			}
			if t == 'R' {
				start = mappings[start][1]
				localCnt++
			}
			if t == 'L' {
				start = mappings[start][0]
				localCnt++
			}
		}
		counts = append(counts, localCnt)
		fmt.Println(key, localCnt)
	}
	if len(counts) == 0 {
		return -1
	}
	if len(counts) == 1 {
		return counts[0]
	}
	return LCM(counts)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(n []int) int {
	result := n[0]
	for i := 1; i < len(n); i++ {
		result = (result * n[i]) / GCD(result, n[i])
	}
	return result
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
