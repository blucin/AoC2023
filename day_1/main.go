package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	solution := 0
	scanner := bufio.NewScanner(input)
	patterns := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for scanner.Scan() {
		line := scanner.Text()
		maxPatternIdx := 0
		minPatternIdx := 99
		min := ""
		max := ""
		for _, pattern := range patterns {
			upperBoundIdx := strings.LastIndex(line, pattern)
			lowerBoundIdx := strings.Index(line, pattern)
			if upperBoundIdx >= maxPatternIdx {
				maxPatternIdx = upperBoundIdx
				letter2Num, exists := numberMap[pattern]
				if exists {
					max = letter2Num
				} else {
					max = pattern
				}
			}
			if lowerBoundIdx <= minPatternIdx && lowerBoundIdx != -1 {
				minPatternIdx = lowerBoundIdx
				letter2Num, exists := numberMap[pattern]
				if exists {
					min = letter2Num
				} else {
					min = pattern
				}
			}
		}

		toNum, err := strconv.Atoi(min + max)
		if err != nil {
			log.Fatal(err)
		}
		solution += toNum
		i++
	}
	fmt.Println(solution)
}
