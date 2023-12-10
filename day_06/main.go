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
	lines := ReadFile("input.txt")
	times := parseLinePuzzle1(lines[0][11:])
	dists := parseLinePuzzle1(lines[1][11:])

	solution1 := 1
	for i, time := range times {
		ways := 0
		for held := 0; held < time; held++ {
			d := (time - held) * held
			if d > 0 && d > dists[i] {
				ways++
			}
		}
		solution1 *= ways
	}
	fmt.Println(solution1)

	time := parseLinePuzzle2(lines[0][11:])
	dist := parseLinePuzzle2(lines[1][11:])
	solution2 := 1
	ways := 0
	for held := 1; held <= time; held++ {
		d := (time - held) * held
		if d > 0 && d > dist {
			ways++
		}
	}
	solution2 *= ways
	fmt.Println(solution2)
	fmt.Println(time - ways)
	fmt.Println(time)
}

// takes a line string num1 num2 num3 num4 ...
// parses a string line into an array
func parseLinePuzzle1(str string) []int {
	result := make([]int, 0)
	for _, s := range strings.Split(str, " ") {
		n := Str2Num(s)
		if n == -1 {
			continue
		}
		result = append(result, n)
	}
	return result
}

func parseLinePuzzle2(str string) int {
	temp := ""
	for _, s := range strings.Split(str, " ") {
		temp += s
	}
	return Str2Num(temp)
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

func Str2Num(s string) int {
	if len(s) == 0 {
		return -1
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
