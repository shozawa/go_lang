package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dup struct {
	line      string
	count     int
	fileNames []string
}

func appendFileName(dup *Dup, name string) {
	if !contains(dup.fileNames, name) {
		dup.fileNames = append(dup.fileNames, name)
	}
}

func contains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func main() {
	counts := make(map[string]Dup)
	files := os.Args[1:]
	if len(files) == 0 {
		counntLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			counntLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, strings.Join(n.fileNames, ","))
		}
	}
}

func counntLines(f *os.File, counts map[string]Dup) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		dup := counts[input.Text()]
		appendFileName(&dup, f.Name())
		dup.count++
		counts[input.Text()] = dup
	}
}
