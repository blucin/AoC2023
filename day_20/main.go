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

const (
	on  = true
	off = false
)

func main() {
	modules, operations := ParseFile(ReadFile("input_test.txt"))
	solution1 := GetProduct(modules, operations)
	fmt.Println("Solution1: ", solution1)
}

func GetProduct(mod map[string]Modules, operations map[string][]string) int {
	convMemories := make(map[string][]bool, 0)
	lowCnt, highCnt := 0, 0
	logs := make([]string, 0)

	initialState := make([]bool, 0)
	for _, v := range mod {
		initialState = append(initialState, v.state)
	}

	// TODO: fix iteration over map being undeterministic
	keys := make([]string, 0)
	for k := range operations {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	iter := 0
	for {
		// exit condition
		currentState := make([]bool, 0)
		for _, v := range mod {
			currentState = append(currentState, v.state)
		}
		if iter > 0 {
			shouldBreak := false
			// flipflops
			for _, state := range currentState {
				if state {
					shouldBreak = true
					break
				}
			}
			// conjunction modules
			for _, convState := range convMemories {
				for _, state := range convState {
					if state {
						shouldBreak = true
						break
					}
				}
			}
			if shouldBreak {
				break
			}
		}

		for _, lhs := range keys {
			rhs := operations[lhs]
			switch {
			case lhs[0] == '%': // flipflop
				for _, val := range rhs {
					if mod[val].prefix == '&' {
						toSend := mod[lhs[1:]].state
						lowCnt, highCnt = GetCount(toSend, lowCnt, highCnt)
						convMemories[val] = append(convMemories[val], toSend)
						logs = append(logs, Logger(lhs[1:], toSend, rhs)...)
						continue
					}
					sent := SwitchFlipFlop(mod[lhs[1:]].state, mod, val)
					mod[val] = Modules{
						prefix: mod[val].prefix,
						state:  sent,
					}
					logs = append(logs, Logger(lhs[1:], sent, rhs)...)
					lowCnt, highCnt = GetCount(sent, lowCnt, highCnt)
				}

			case lhs[0] == '&': // conjuction
				toSend := off
				for _, val := range convMemories[lhs[1:]] {
					if !val {
						toSend = on
						break
					}
				}
				for _, val := range rhs {
					SwitchFlipFlop(toSend, mod, val)
					mod[val] = Modules{
						prefix: mod[val].prefix,
						state:  toSend,
					}
					logs = append(logs, Logger(lhs[1:], toSend, rhs)...)
					lowCnt, highCnt = GetCount(toSend, lowCnt, highCnt)
				}

			case lhs == "broadcaster":
				for _, val := range rhs {
					mod[val] = Modules{
						prefix: mod[val].prefix,
						state:  off,
					}
					lowCnt++
				}

			default:
				log.Fatal("Invalid operation found")
			}
		}
		iter++
	}
	fmt.Println("Low: ", lowCnt, "High: ", highCnt)
	for _, log := range logs {
		fmt.Println(log)
	}
	return lowCnt * highCnt
}

// logger for debugging
func Logger(lhs string, signal bool, rhs []string) []string {
	logs := make([]string, 0)
	for _, sents := range rhs {
		// example: a -high-> b
		signalStr := "low"
		if signal {
			signalStr = "high"
		}
		logs = append(logs, fmt.Sprintf("%s -%s-> %s", lhs, signalStr, sents))
	}
	return logs
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

func ParseFile(lines []string) (map[string]Modules, map[string][]string) {
	operations := make(map[string][]string, 0)
	modules := make(map[string]Modules, 0)
	for _, line := range lines {
		split := strings.Split(line, "->")
		sender := strings.Fields(split[0])[0][0:]
		receiver := strings.FieldsFunc(split[1], func(r rune) bool {
			return r == ',' || r == ' '
		})
		operations[sender] = receiver
		if sender != "broadcaster" {
			pre := rune(strings.Fields(split[0])[0][0])
			name := strings.Fields(split[0])[0][1:]
			modules[name] = Modules{
				prefix: pre,
				state:  off,
			}
		}
	}
	return modules, operations
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
