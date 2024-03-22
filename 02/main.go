package main

import (
	"fmt"
	"os"
	"time"
)

type Rounds struct {
	round [][]byte
}

func main() {
	fileName := "test.txt"
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	rounds := insertData(file)
	s1 := time.Now()

}

func insertData(file []byte) Rounds {
	data := strings.Split(string(file), "\n\r\n")
	for _, datum := range data {

	}
}
