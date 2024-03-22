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
	fmt.Println(result1)
	fmt.Println("Time in nanoseconds:", time.Since(s1).Nanoseconds())

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
		fmt.Println(m[round[1]], m[round[0]+round[1]])
		result += m[round[0]+round[1]]

	}
	return result
}
