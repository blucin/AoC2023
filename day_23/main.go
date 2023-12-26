package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := ReadFile("input_test.txt")
	solution1 := FindLongestPath(lines)
	fmt.Println("Solution 1:", solution1)
}

type Node struct {
	r, c int // row and column number
	n    int // distance from the source
}

type Stack []*Node

func New() *Stack {
	return &Stack{}
}
func (s *Stack) Len() int     { return len(*s) }
func (s *Stack) Peek() *Node  { return (*s)[s.Len()-1] }
func (s *Stack) Push(n *Node) { *s = append(*s, n) }
func (s *Stack) Pop() *Node   { n := s.Peek(); *s = (*s)[:s.Len()-1]; return n }

func DFS(node *Node, matrix [][]rune, seen map[[2]int]bool, currentDist int, maxDist *int) {
	seen[[2]int{node.r, node.c}] = true
	for _, neighbour := range GetNeighbours(node.r, node.c, matrix) {
		if !seen[[2]int{neighbour[0], neighbour[1]}] {
			DFS(&Node{neighbour[0], neighbour[1], node.n + 1}, matrix, seen, currentDist+1, maxDist)
		}
	}
	if currentDist > *maxDist {
		*maxDist = currentDist
	}
	seen[[2]int{node.r, node.c}] = false
}

func FindLongestPath(matrix [][]rune) int {
	seen := make(map[[2]int]bool)
	maxDist := -1
	DFS(&Node{0, 1, 0}, matrix, seen, 0, &maxDist) // start node
	return maxDist
}

func isValidMove(r int, c int, matrix [][]rune) bool {
	if r > 0 && r < len(matrix) && c > 0 && c < len(matrix[0]) && matrix[r][c] != '#' {
		return true
	}
	return false
}

func GetNeighbours(x int, y int, matrix [][]rune) [][2]int {
	switch matrix[x][y] {
	case '^':
		return [][2]int{{x - 1, y}}
	case 'v':
		return [][2]int{{x + 1, y}}
	case '<':
		return [][2]int{{x, y - 1}}
	case '>':
		return [][2]int{{x, y + 1}}
	default:
		neighbours := make([][2]int, 0)
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range directions {
			if !isValidMove(x+dir[0], y+dir[1], matrix) {
				continue
			}
			neighbours = append(neighbours, [2]int{x + dir[0], y + dir[1]})
		}
		return neighbours
	}
}

func ReadFile(fileName string) [][]rune {
	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	lines := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	return lines
}
