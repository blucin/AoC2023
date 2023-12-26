package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	on  = true
	off = false
)

func main() {
	broadcasts, modules, operations := ParseFile(ReadFile("input_test_test.txt"))
	solution1 := GetProduct(broadcasts, modules, operations)
	fmt.Println("Solution1: ", solution1)

}

func GetProduct(broadcasts []string, mod map[string]Modules, operations map[string][]string) int {
	convMemories := make(map[string][]bool, 0)
	lowCnt, highCnt := 0, 0

	initialState := make([]bool, 0)
	for _, v := range mod {
		initialState = append(initialState, v.state)
	}

	iter := 0
	for {
		// exit condition
		currentState := make([]bool, 0)
		for _, v := range mod {
			currentState = append(currentState, v.state)
		}
		if CompareSlices(initialState, currentState) && iter > 0 {
			break
		}

		for lhs, rhs := range operations {
			switch {
			case lhs[0] == '%': // flipflop
				for _, val := range rhs {
					if mod[val].prefix == '&' {
						toSend := mod[lhs[1:]].state
						lowCnt, highCnt = GetCount(toSend, lowCnt, highCnt)
						convMemories[val] = append(convMemories[val], toSend)
						continue
					}
					sent := SwitchFlipFlop(mod[lhs[1:]].state, mod, val)
					lowCnt, highCnt = GetCount(sent, lowCnt, highCnt)
				}

			case lhs[0] == '&': // conjuction
				convState := GetConjunctionState(convMemories[lhs[1:]])
				for _, val := range rhs {
					SwitchFlipFlop(convState, mod, val)
					lowCnt, highCnt = GetCount(convState, lowCnt, highCnt)
				}

			case lhs[0] == 'b': // broadcast
				lowCnt += BroadcastOffSignal(mod, broadcasts)

			default:
				log.Fatal("Invalid operation found")
			}
		}
		iter++
	}
	fmt.Println("Low: ", lowCnt, "High: ", highCnt)
	return lowCnt * highCnt
}

type Modules struct {
	prefix rune
	state  bool
}

// calculates count of low and high signal
func GetCount(state bool, lowCnt int, highCnt int) (int, int) {
	if state {
		return lowCnt, highCnt + 1
	}
	return lowCnt + 1, highCnt
}

// if one of the signal is low, then send high
// if all signals in the memory are high, then send low
func GetConjunctionState(memory []bool) bool {
	for _, val := range memory {
		if !val {
			return on
		}
	}
	return off
}

// broadcasts low signal and returns
// number of low signals broadcasted
func BroadcastOffSignal(mod map[string]Modules, names []string) int {
	lowCnt := 0
	for _, name := range names {
		mod[name] = Modules{
			prefix: mod[name].prefix,
			state:  off,
		}
		lowCnt++
	}
	return lowCnt
}

func CompareSlices(s1 []bool, s2 []bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, val := range s1 {
		if val != s2[i] {
			return false
		}
	}
	return true
}

// if signal is high, then do nothing
// if signal is low, then switch state
func SwitchFlipFlop(signal bool, mod map[string]Modules, name string) bool {
	if signal {
		return signal
	}
	if mod[name].state {
		mod[name] = Modules{
			prefix: mod[name].prefix,
			state:  off,
		}
		return off
	} else {
		mod[name] = Modules{
			prefix: mod[name].prefix,
			state:  on,
		}
		return on
	}
}

func ParseFile(lines []string) ([]string, map[string]Modules, map[string][]string) {
	broadcasts := make([]string, 0)
	operations := make(map[string][]string, 0)
	modules := make(map[string]Modules, 0)
	for i, line := range lines {
		s1 := strings.Fields(line)
		if strings.HasPrefix(line, "broadcaster") { // broadcaster
			for j := 2; j < len(s1); j++ {
				withoutComma := strings.TrimSuffix(s1[j], ",")
				broadcasts = append(broadcasts, withoutComma)
			}
		}
		if i > 0 { // operations and modules
			split := strings.Split(line, "->")
			sender := strings.Fields(split[0])[0][0:]
			receiver := strings.FieldsFunc(split[1], func(r rune) bool {
				return r == ',' || r == ' '
			})
			operations[sender] = receiver
			pre := rune(strings.Fields(split[0])[0][0])
			name := strings.Fields(split[0])[0][1:]
			modules[name] = Modules{
				prefix: pre,
				state:  off,
			}
		}
	}
	return broadcasts, modules, operations
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
