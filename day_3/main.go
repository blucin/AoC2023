package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// problem constraint
func Check(r rune) bool {
	return r == '*' || r == '#' || r == '+' || r == '$'
}

func Str2Num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

var match = map[rune]bool{
	'=': true,
	'*': true,
	'#': true,
	'$': true,
	'@': true,
	'&': true,
	'/': true,
	'-': true,
	'+': true,
	'%': true,
}

var lines []string

func isNum(r rune) bool {
	return unicode.IsDigit(r)
}

func CheckValidChar(s string) bool {
	for _, c := range s {
		_, exists := match[c]
		if exists {
			return true
		}
	}
	return false
}

func main() {
	input, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	solution1 := uint(0)
	solution2 := uint(0)

	scanner := bufio.NewScanner(input)

	// load file in lines[]
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for i, line := range lines {
		stack := ""
		lineSol := 0
		for j, char := range line {
			if unicode.IsDigit(char) {
				stack += string(char)
			}
			if (!unicode.IsDigit(char) && stack != "") || (j == len(line)-1 && stack != "") {
				if j == len(line)-1 && unicode.IsDigit(rune(line[j])) { // edge case last char
					j++
				}
				checkStr := ""
				first := j - len(stack) // first digit
				last := j               // last digit
				if i > 0 {
					checkStr += string(lines[i-1][first:last]) // above
					if first > 0 {
						checkStr += string(lines[i][first-1])   // left
						checkStr += string(lines[i-1][first-1]) // above diagonal left
					}
					if last < len(line)-1 {
						checkStr += string(lines[i-1][last]) // above diagonal right
					}
				}
				if i < len(lines)-1 {
					checkStr += string(lines[i+1][first:last]) // below
					if first > 0 {
						checkStr += string(lines[i+1][first-1]) // below diagonal  left
					}
					if last < len(line)-1 {
						checkStr += string(lines[i][last])   // right
						checkStr += string(lines[i+1][last]) // below diagonal  right
					}
				}
				if CheckValidChar(checkStr) {
					lineSol += Str2Num(stack)
				}
				stack = ""
			}
		}
		solution1 += uint(lineSol)
	}

	// Part 1: 925
	fmt.Printf("Solution of puzzle1: %d\n", solution1)
	fmt.Printf("Solution of puzzle2: %d\n", solution2)
}
