package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Str2Num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// puzzle 1
func calculateLineSol(winningNums []uint, numsFound []uint) int {
	lineSol := 0
	cnt := 0
	for _, num := range numsFound {
		for _, winNum := range winningNums {
			if num == winNum {
				if cnt == 0 {
					lineSol = 1
				} else {
					lineSol *= 2
				}
				cnt++
			}
		}
	}
	return lineSol
}

// puzzle 2
func getCardMatchIdx(winningNums []uint, numsFound []uint, cardIdx uint) []uint {
	cardMatches := make([]uint, 0)
	matchCnt := 0
	for _, num := range numsFound {
		for _, winNum := range winningNums {
			if num == winNum {
				matchCnt++
				cardMatches = append(cardMatches, uint(matchCnt+int(cardIdx)))
			}
		}
	}
	return cardMatches
}

var lines []string

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

	cardMatchCnt := make(map[uint]uint) // puzzle2

	for i, line := range lines {
		lineSol := 0
		stack := ""
		winningNums := make([]uint, 0)
		numsFound := make([]uint, 0)
		barFound := false
		line = line[8:]
		for j, char := range line {
			if char == '|' {
				barFound = true
			}
			if unicode.IsDigit(char) {
				stack += string(char)
			}
			if !unicode.IsDigit(char) && !barFound && stack != "" {
				num := Str2Num(stack)
				winningNums = append(winningNums, uint(num))
				stack = ""
			}
			if !unicode.IsDigit(char) && barFound && stack != "" || (j == len(line)-1 && stack != "") {
				num := Str2Num(stack)
				numsFound = append(numsFound, uint(num))
				stack = ""
			}
		}
		lineSol = calculateLineSol(winningNums, numsFound)

		// puzzle 2
		cardIdx := uint(i) + 1
		lineMatches := getCardMatchIdx(winningNums, numsFound, cardIdx)
		cardMatchCnt[cardIdx] += 1
		copies := cardMatchCnt[cardIdx]
		for _, match := range lineMatches {
			cardMatchCnt[match] += copies
		}

		solution1 += uint(lineSol)
	}

	// solution2 is the sum of all the values from key-values pairs of cardMatchCnt
	for _, v := range cardMatchCnt {
		solution2 += v
	}

	fmt.Printf("Solution of puzzle1: %d\n", solution1)
	fmt.Printf("Solution of puzzle2: %d\n", solution2)
}
