package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Line struct {
	hand string
	bid  int
}

var p1labelRanks = map[rune]int{'2': 2, '3': 3, '4': 4,
	'5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12,
	'K': 13, 'A': 14,
}

// in puzzle2 J is give the lowest rank
var p2labelRanks = map[rune]int{'J': 1, '2': 2, '3': 3, '4': 4,
	'5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'Q': 11, 'K': 12,
	'A': 13,
}

var turnPuzzle2On = false

func main() {
	lines := parseLines(readFile("input.txt"))

	// puzzle 1
	solution1 := 0
	lines = sortLines(lines, p1labelRanks)
	for i, line := range lines {
		solution1 += line.bid * (i + 1)
	}
	fmt.Println("solution 1:", solution1)

	// puzzle 2
	turnPuzzle2On = true
	solution2 := 0
	lines = sortLines(lines, p2labelRanks)
	for i, line := range lines {
		solution2 += line.bid * (i + 1)
	}
	fmt.Println("solution 2:", solution2)
}

func sortLines(lines []Line, labelRanks map[rune]int) []Line {
	sort.Slice(lines, func(i, j int) bool {
		if rank(lines[i].hand) == rank(lines[j].hand) {
			for labelIdx, label := range lines[i].hand {
				if labelRanks[label] > labelRanks[rune(lines[j].hand[labelIdx])] {
					return false
				}
				if labelRanks[label] < labelRanks[rune(lines[j].hand[labelIdx])] {
					return true
				}
			}
			return false
		}
		return rank(lines[i].hand) < rank(lines[j].hand)
	})
	return lines
}

func rank(hand string) int {
	cnt := make(map[rune]int)
	for _, c := range hand {
		cnt[c]++
	}
	if turnPuzzle2On {
		_, exists := cnt['J']
		if exists && hand != "JJJJJ" {
			// transfer J's cnt to the key with max value
			temp := 0
			maxKey := '0'
			for k, v := range cnt {
				if v > temp && k != 'J' {
					maxKey = k
					temp = v
				}
			}
			cnt[maxKey] += cnt['J']
			delete(cnt, 'J')
		}
	}
	switch len(cnt) {
	case 1:
		return 7 // 5 of a kind
	case 2:
		if valueExists(cnt, 4) {
			return 6 // 4 of a kind
		} else {
			return 5 // full house
		}
	case 3:
		if valueExists(cnt, 3) {
			return 4 // 3 of a kind
		} else {
			return 3 // 2 pair
		}
	case 4: // 1 pair
		return 2
	case 5: // high card
		return 1
	default:
		log.Fatal("invalid hand")
	}
	return -1
}

func valueExists(cnt map[rune]int, val int) bool {
	for _, v := range cnt {
		if v == val {
			return true
		}
	}
	return false
}

func parseLines(lines []string) []Line {
	parsed := make([]Line, 0)
	for _, line := range lines {
		l, r := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		parsed = append(parsed, Line{hand: l, bid: Str2Num(r)})
	}
	return parsed
}

func Str2Num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func readFile(fileName string) []string {
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
