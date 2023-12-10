package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := parseLines(ReadFile("input_test.txt"))
	solution1 := 0
	solution2 := 0
	for _, line1 := range lines {
		triangle := make([][]int, 0)
		triangle = append(triangle, line1)
		for checkZeroes(triangle[len(triangle)-1]) {
			next := make([]int, len(triangle[len(triangle)-1])-1)
			for i := range triangle[len(triangle)-1] {
				if i < len(triangle[len(triangle)-1])-1 {
					next[i] = triangle[len(triangle)-1][i+1] - triangle[len(triangle)-1][i]
				}
			}
			triangle = append(triangle, next)
		}
		curr := 0
		for i := len(triangle) - 1; i >= 0; i-- {
			solution1 += triangle[i][len(triangle[i])-1]
			curr = triangle[i][0] - curr
		}
		solution2 += curr
	}
	fmt.Println("Solution 1:", solution1)
	fmt.Println("Solution 2:", solution2)
}

func checkZeroes(line []int) bool {
	for _, v := range line {
		if v != 0 {
			return true
		}
	}
	return false
}

func parseLines(lines []string) [][]int {
	n := make([][]int, 0)
	for _, line := range lines {
		s := strings.Split(line, " ")
		arr := make([]int, len(s))
		for i, v := range s {
			arr[i] = Str2Num(v)
		}
		n = append(n, arr)
	}
	return n
}

func Str2Num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
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
