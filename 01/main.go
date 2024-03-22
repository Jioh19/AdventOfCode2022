package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fileName := "input.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	s1 := time.Now()
	array := insertData(file)
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
	fmt.Println("First int: ", array[0])
	fmt.Println("Time in nanoseconds:", time.Since(s1).Nanoseconds())

	s2 := time.Now()
	fmt.Println("Total Calories: ", array[0]+array[1]+array[2])
	fmt.Println("Time in nanoseconds:", time.Since(s2).Nanoseconds())
}

func insertData(file []byte) []int {
	var array []int
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	counter := 0
	for _, line := range lines {
		if line != "" {
			aux, _ := strconv.Atoi(line)
			counter += aux
		} else {
			array = append(array, counter)
			counter = 0
		}
	}
	return array
}
