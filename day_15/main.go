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
	// input
	line := ReadFile("input_test.txt")[0]
	solution1, solution2 := InitBoxAndGetHashSum(line)
	fmt.Println("solution 1:", solution1)
	fmt.Println("solution 2:", solution2)
}

func InitBoxAndGetHashSum(input string) (hashSum int, totalFocus int) {
	boxes := make(map[int][][2]any, 0)

SkipLetter:
	for _, letters := range strings.Split(input, ",") {
		// puzzle 1
		hashSum += Hash(letters)

		// puzzle 2
		if letters[len(letters)-1] == '-' {
			label := strings.Split(letters, "-")[0]
			for i, arr := range boxes[Hash(label)] {
				if arr[0] == label {
					boxes[Hash(label)] = remove(boxes[Hash(label)], i)
					continue SkipLetter
				}
			}
			continue SkipLetter
		}

		split := strings.Split(letters, "=")
		boxKey, label, lens := Hash(split[0]), split[0], split[1]
		for i, arr := range boxes[boxKey] {
			// "=" : condition 1
			if arr[0] == label {
				boxes[boxKey][i] = [2]any{label, lens}
				continue SkipLetter
			}
		}
		// "=" : condition 2
		boxes[boxKey] = append(boxes[boxKey], [2]any{label, lens})
	}

	for boxNum, box := range boxes {
		for slotNum, slot := range box {
			totalFocus += (1 + boxNum) * (slotNum + 1) * Str2Num(slot[1].(string))
		}
	}
	return hashSum, totalFocus
}

func remove[T any](s []T, i int) []T {
	// order it not important
	// s[i] = s[len(s)-1]
	// return s[:len(s)-1]

	// order is important
	return append(s[:i], s[i+1:]...)
}

func Hash(s string) int {
	hash := 0
	for _, c := range s {
		hash = ((hash + int(c)) * 17) % 256
	}
	return hash
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
