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
	lines := ReadFile("input_test.txt")

	// puzzle1
	workflows, ratings := ParseFile(lines)
	solution1 := FindAcceptedRatingsSum(workflows, ratings)
	fmt.Println("Solution 1:", solution1)

	// puzzle2: partMap consists of all possible workflow from 1 to 4k
}

func FindAcceptedRatingsSum(workflows map[string]string, ratings [][]int) int {
	accepted := 0
	partMap := make(map[rune]int, 0)
	for _, rating := range ratings {
		partMap = map[rune]int{
			'x': rating[0],
			'm': rating[1],
			'a': rating[2],
			's': rating[3],
		}
		startingExprs := strings.Split(workflows["in"], ",")
		if IsValidRating(workflows, startingExprs, partMap) {
			for _, v := range rating {
				accepted += v
			}
		}
	}
	return accepted
}

func IsValidRating(workflows map[string]string, startExprs []string, partMap map[rune]int) bool {
	for _, expr := range startExprs {
		// exit condition
		if len(expr) == 1 {
			return expr == "A"
		}

		// check if the expression is only a key
		if expr[1] != '>' && expr[1] != '<' {
			nextKey := expr
			nextExprs := strings.Split(workflows[nextKey], ",")
			return IsValidRating(workflows, nextExprs, partMap)
		}

		partValue, numValue := partMap[rune(expr[0])], Str2Num(strings.Split(expr, ":")[0][2:])
		isValidExpr := IsValidExpression(rune(expr[1]), partValue, numValue)

		if !isValidExpr {
			continue
		}

		split := strings.Split(expr, ":")

		if len(split[1]) == 1 {
			return split[1] == "A"
		}

		nextKey := split[1]
		nextExprs := strings.Split(workflows[nextKey], ",")
		return IsValidRating(workflows, nextExprs, partMap)
	}
	return false
}

func IsValidExpression(operator rune, partValue, numValue int) bool {
	switch operator {
	case '>':
		return partValue > numValue
	case '<':
		return partValue < numValue
	default:
		return false
	}
}

func ParseFile(lines []string) (map[string]string, [][]int) {
	workflows := make(map[string]string, 0)
	ratings := make([][]int, 0)
	flag := false
	for _, line := range lines {
		if line == "" {
			flag = true
			continue
		}
		if flag {
			split := strings.Split(line[1:len(line)-1], ",")
			rating := make([]int, 0)
			for _, num := range split {
				rating = append(rating, Str2Num(strings.Split(num, "=")[1]))
			}
			ratings = append(ratings, rating)
		} else {
			split := strings.Split(line, "{")
			workflows[split[0]] = split[1][0 : len(split[1])-1]
		}
	}
	return workflows, ratings
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
