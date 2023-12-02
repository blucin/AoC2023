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
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	solution1 := uint(0)
	solution2 := uint(0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lineCnt := map[string]uint{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		GetMaxLineCount(line, lineCnt)
		if lineCnt["red"] <= 12 && lineCnt["green"] <= 13 && lineCnt["blue"] <= 14 {
			solution1 += uint(i)
		}
		solution2 += lineCnt["red"] * lineCnt["green"] * lineCnt["blue"]
		i++
	}
	fmt.Printf("Solution of puzzle1: %d\n", solution1)
	fmt.Printf("Solution of puzzle2: %d\n", solution2)
}

func SplitSubGames(r rune) bool {
	return r == ' ' || r == ','
}

func SplitGame(r rune) bool {
	return r == ':' || r == ';'
}

func GetMaxLineCount(line string, cnt map[string]uint) {
	splitGame := strings.FieldsFunc(line, SplitGame)
	for i, str := range splitGame {
		subGame := strings.FieldsFunc(str, SplitSubGames)
		if i == 0 {
			continue
		} // ignore Game Num label

		for j := 0; j < len(subGame); j += 2 {
			num, err := strconv.Atoi(subGame[j])
			if err != nil {
				log.Fatal(err)
			}
			if cnt[subGame[j+1]] < uint(num) {
				cnt[subGame[j+1]] = uint(num)
			}
		}
	}
}
