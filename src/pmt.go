package main

import "github.com/oknotok97/pmt/src/cmd"

func main() {
	cmd.BuildCli()
	//file, err := os.Open("shakespeare.txt")
	//if err != nil {
	//	log.Fatalf("Could not open file")
	//}
	//defer file.Close()
	//scanner := bufio.NewScanner(file)
	//lineCount := 0
	//pat := "the"
	//var occ []int
	//sw := stopwatch.CreateStarted()
	//
	//for scanner.Scan() {
	//	occ = shiftor.Search(scanner.Text(), pat)
	//	if len(occ) > 0 {
	//		lineCount += 1
	//	}
	//}
	//sw.Stop()
	//fmt.Printf("Line Count: %d\n in %.2fms", lineCount, sw.GetElapsedTime())
}
