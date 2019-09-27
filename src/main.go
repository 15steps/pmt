package main

import (
	"bufio"
	"fmt"
	"github.com/oknotok97/pmt/src/shiftor"
	"github.com/oknotok97/stopwatch"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./data/shakespeare.txt")
	if err != nil {
		log.Fatalf("Could not open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	pat := "Romeo"
	var occ []int
	sw := stopwatch.CreateStarted()

	for scanner.Scan() {
		occ = shiftor.Search(scanner.Text(), pat)
		if len(occ) > 0 {
			lineCount += len(occ)
		}
	}
	sw.Stop()
	fmt.Printf("Line Count: %d\n in %.2fms", lineCount, sw.GetElapsedTime())
}
