package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction rune
	distance  int
}

func main() {
	instructions := ParseFilePuzzle1(ReadFile("input_test.txt"))
	fmt.Println("Solution1: ", FindArea(instructions))

	instructions = ParseFilePuzzle2(ReadFile("input_test.txt"))
	fmt.Println("Solution2: ", FindArea(instructions))
}

func FindArea(instrs []Instruction) int {
	area := 0
	boundaryArea := 0
	vertices := make([][2]int, 0)
	start := [2]int{0, 0}
	vertices = append(vertices, start)

	for _, instr := range instrs {
		prev := vertices[len(vertices)-1]
		switch instr.direction {
		case 'R':
			vertices = append(vertices, [2]int{prev[0] + instr.distance, prev[1]})
		case 'L':
			vertices = append(vertices, [2]int{prev[0] - instr.distance, prev[1]})
		case 'U':
			vertices = append(vertices, [2]int{prev[0], prev[1] + instr.distance})
		case 'D':
			vertices = append(vertices, [2]int{prev[0], prev[1] - instr.distance})
		}
		boundaryArea += instr.distance
	}

	// shoelace formula
	for i := range vertices {
		x1 := vertices[i][0]
		y1 := vertices[i][1]
		x2 := vertices[(i+1)%len(vertices)][0] // avoid out of bounds
		y2 := vertices[(i+1)%len(vertices)][1] // avoid out of bounds
		area += (x1 * y2) - (x2 * y1)
	}

	return Abs(area)/2 + (boundaryArea / 2) + 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ParseFilePuzzle2(lines []string) []Instruction {
	instrs := make([]Instruction, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		clr := split[2][2 : len(split[2])-1]
		dir := 'X'
		switch clr[len(clr)-1] {
		case '0':
			dir = 'R'
		case '1':
			dir = 'D'
		case '2':
			dir = 'L'
		case '3':
			dir = 'U'
		}
		dist := new(big.Int)
		dist.SetString(split[2][2:len(split[2])-2], 16)
		instrs = append(instrs, Instruction{dir, int(dist.Int64())})
	}
	return instrs
}

func ParseFilePuzzle1(lines []string) []Instruction {
	instrs := make([]Instruction, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		dir := rune(split[0][0])
		dist := Str2Num(split[1])
		instrs = append(instrs, Instruction{dir, dist})
	}
	return instrs
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
