package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

func main() {
	grid := ParseGrid(ReadFile("input_test.txt"))
	solution1 := GetShortestHeatLoss(grid)
	fmt.Println("Solution 1:", solution1+1)
}

type Node struct {
	r, c   int // row, col
	thl    int // total heat loss
	dr, dc int // direction
	cnt    int // cnt of prev node with same dr, dc
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

// We want pop to give lowest
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].thl < pq[j].thl
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func GetShortestHeatLoss(grid [][]int) int {
	minHl := make([]Node, 0)
	src := Node{0, 0, 0, 0, 0, 0} // starting node
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	visualDebug := make([][]int, len(grid))
	for i := range visualDebug {
		visualDebug[i] = make([]int, len(grid[i]))
	}

	// push the initial node with 0 hl
	pq.Push(&src)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		visualDebug[node.r][node.c] = node.thl

		// reached the last node (exit condition return it's hl from the starting node)
		if node.c == len(grid[0])-1 && node.r == len(grid)-1 {
			return node.thl
		}

		// check if the node is already in minHl (it might not have the same hl)
		found := false
		for _, nde := range minHl {
			if nde.r == node.r && nde.c == node.c && nde.dr == node.dr && nde.dc == node.dc && node.cnt == nde.cnt {
				found = true
				break
			}
		}
		if found {
			continue
		}
		minHl = append(minHl, *node)

		// we are free to move in a straight path for not more than 3 times
		if node.cnt < 3 && !(node.dr == 0 && node.dc == 0) {
			nr := node.r + node.dr
			nc := node.c + node.dc
			if nr >= 0 && nr <= len(grid)-1 && nc >= 0 && nc <= len(grid[0])-1 {
				next := &Node{nr, nc, node.thl + grid[nr][nc], node.dr, node.dc, node.cnt + 1}
				heap.Push(&pq, next)
				visualDebug[nr][nc] = node.thl + grid[nr][nc]
			}
		}

		// if the above one can't find the shortest path or we can't move in the same line anymore
		// move to other directions
		for _, nd := range dir {
			if (nd[0] == -node.dr && nd[0] == -node.dc) || (nd[0] == node.dr && nd[1] == node.dc) {
				continue
			}
			nr := node.r + nd[0]
			nc := node.c + nd[1]
			if nr >= 0 && nr <= len(grid)-1 && nc >= 0 && nc <= len(grid[0])-1 {
				next := &Node{nr, nc, node.thl + grid[nr][nc], nd[0], nd[1], 1}
				heap.Push(&pq, next)
				visualDebug[nr][nc] = node.thl + grid[nr][nc]
			}
		}
	}
	// no route with min hl found from the start to end
	return -1
}

func ParseGrid(grid [][]rune) [][]int {
	parsed := make([][]int, len(grid))
	for i := range grid {
		parsed[i] = make([]int, len(grid[i]))
		for j := range grid[i] {
			parsed[i][j] = int(grid[i][j] - '0')
		}
	}
	return parsed
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
