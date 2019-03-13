package main

import (
	"fmt"
	"strings"
	"time"
)

func naiveJoin(lines []string) {
	start := time.Now()
	s, sep := "", ""
	for _, line := range lines {
		s += sep + line
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start)
	fmt.Printf("naiveJoin time\t%d\n", secs)
}

func smartJoin(lines []string) {
	start := time.Now()
	fmt.Println(strings.Join(lines, " "))
	secs := time.Since(start)
	fmt.Printf("smartJoin time\t%d\n", secs)
}

func main() {
	lines := []string{"aaaaaaa", "bbbbb"}
	naiveJoin(lines)
	smartJoin(lines)
}
