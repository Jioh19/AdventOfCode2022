package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Rounds struct {
	round [][]string
}

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := make(map[string]int)
	m["AX"] = 3
	m["AY"] = 6
	m["BY"] = 3
	m["BZ"] = 6
	m["CZ"] = 3
	m["CX"] = 6
	m["X"] = 1
	m["Y"] = 2
	m["Z"] = 3
	s1 := time.Now()
	rounds := insertData(file)
	result1 := countPoints(rounds, m)
	fmt.Println("Part 1: ", result1)
	fmt.Println("Time in nanoseconds:", time.Since(s1).Nanoseconds())

	s2 := time.Now()
	result2 := countPoints2(rounds)
	fmt.Println("Part 2: ", result2)
	fmt.Println("Time in nanoseconds:", time.Since(s2).Nanoseconds())

}

func insertData(file []byte) Rounds {
	var result Rounds
	data := strings.Split(string(file), "\n")
	for _, datum := range data {
		round := strings.Split(datum, " ")
		result.round = append(result.round, round)
	}
	return result
}

func countPoints(rounds Rounds, m map[string]int) int {
	result := 0
	for _, round := range rounds.round {
		result += m[round[1]]
		result += m[round[0]+round[1]]
	}
	return result
}

func countPoints2(rounds Rounds) int {
	result := 0

	m := make(map[string]int)
	m["AX"] = 3
	m["AY"] = 1
	m["AZ"] = 2
	m["BY"] = 2
	m["BZ"] = 3
	m["BX"] = 1
	m["CZ"] = 1
	m["CX"] = 2
	m["CY"] = 3
	m["X"] = 0
	m["Y"] = 3
	m["Z"] = 6
	for _, round := range rounds.round {
		result += m[round[1]]
		result += m[round[0]+round[1]]
	}
	return result
}
