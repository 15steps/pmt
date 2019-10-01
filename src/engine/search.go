package engine

import (
	"bufio"
	"fmt"
)

// algorithms
const (
	ShiftOr = "shiftor"
)

const Red = "\033[91m"
const EndColor = "\033[0m"

type SearchEngine interface {
	Search(txt string, pat string) []int
}

var engines = map[string]SearchEngine{
	ShiftOr: NewShiftOr(),
}

func SearchHandler(pattern string, fileNames []string, algorithm string) error {
	if algorithm == "" {
		algorithm = ShiftOr
	}
	engine := engines[algorithm]

	files, err := GetFiles(fileNames)
	if err != nil {
		return err
	}

	for _, file := range files {
		//fmt.Printf("-----------------> %s\n", fileNames[i])
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		searchFile(scanner, pattern, engine)
	}

	return nil
}

func searchFile(scanner *bufio.Scanner, pattern string, engine SearchEngine) {
	var occ []int
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		occ = engine.Search(text, pattern)
		if len(occ) > 0 {
			printOcc(text, pattern, occ)
			count +=1
		}
	}
	fmt.Printf("line count: %d\n", count)
}

func printOcc(text string, pat string, occ []int) {
	j := 0
	for i := 0; i < len(text); i++ {
		if j < len(occ) && i == occ[j] {
			colorPrint(pat)
			j++
			i += len(pat) -1
		} else {
			fmt.Printf("%c", text[i])
		}
	}
	fmt.Println()
}

func colorPrint(txt string) {
	fmt.Printf("%s%s%s", Red, txt, EndColor)
}