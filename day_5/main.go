package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type SubMap struct {
	rangeMaps [][]uint
}

// aliases map string string
var A = map[string]string{
	"ss": "seed-to-soil",
	"sf": "soil-to-fertilizer",
	"fw": "fertilizer-to-water",
	"wl": "water-to-light",
	"lt": "light-to-temperature",
	"th": "temperature-to-humidity",
	"hl": "humidity-to-location",
}

func main() {
	lines := readFile("input.txt")
	initialSeeds := make([]uint, 0)
	for _, str := range strings.Split(lines[0][7:], " ") {
		initialSeeds = append(initialSeeds, Str2Num(str))
	}
	ac := parseMap(lines)
	minLocation := findMinLocation(initialSeeds, ac)
	fmt.Printf("solution1: %d\n", minLocation)
}

func findMinLocation(initialSeeds []uint, ac map[string]SubMap) uint {
	minLocation := ^uint(0)
	for _, seed := range initialSeeds {
		temp := get(get(get(get(get(get(get(seed, A["ss"], ac), A["sf"], ac), A["fw"], ac), A["wl"], ac), A["lt"], ac), A["th"], ac), A["hl"], ac)
		if temp < minLocation {
			minLocation = temp
		}
	}
	return minLocation
}

func get(seed uint, mapName string, almanac map[string]SubMap) uint {
	// [value, key, range/width]
	subMap := almanac[mapName]
	match := seed // no match -> seed itself
	for _, rangeMap := range subMap.rangeMaps {
		offset := seed - rangeMap[1]
		if offset < 0 {
			continue
		} // no match
		if offset > rangeMap[2] {
			continue
		} // no match
		match = rangeMap[0] + offset
	}
	return match
}

// returns a map of map name to submap
// submap contains a list of rangemap
// rangemap is a list of uint
// each list has following format: [value, key, range/width]
func parseMap(lines []string) map[string]SubMap {
	almanac := make(map[string]SubMap, 0)
	mapName := ""
	submap := SubMap{rangeMaps: make([][]uint, 0)}
	for _, line := range lines[2:] {
		switch {
		case len(line) == 0: // empty space
			// do nothing
		case unicode.IsLetter(rune(line[0])): // map name
			mapName = strings.Split(line, " ")[0]
			submap.rangeMaps = make([][]uint, 0)
		case unicode.IsDigit(rune(line[0])): // map data
			parsedRangeMap := make([]uint, 0)
			for _, str := range strings.Split(line, " ") {
				parsedRangeMap = append(parsedRangeMap, Str2Num(str))
			}
			submap.rangeMaps = append(submap.rangeMaps, parsedRangeMap)
			almanac[mapName] = submap
		}
	}
	return almanac
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

func Str2Num(s string) uint {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return uint(n)
}
