package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// os is package used for command line arguments access
	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) == 0 {
		// os is the package that has access to Standard IO/Error through files
		// pointing to named Stdin Stdout Stderr
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			// It is equivalent to java: new File(filepath)
			// os is the package used for File IO
			file, err := os.Open(filename)
			if err != nil {
				fmt.Printf("%v\n", os.Stderr)
				continue
			} else {
				countLines(file, counts)
				file.Close()
			}
		}
	}
	printDuplicates(counts)
}

func countLines(file *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			counts[line]++
		} else {
			break
		}
	}
}

func printDuplicates(counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s\t%d\n", line, count)
		}
	}
}
