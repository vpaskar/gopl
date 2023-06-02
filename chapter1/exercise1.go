package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func exercise1() {
	var result, sep string
	for _, arg := range os.Args {
		result += sep + arg
		sep = " "
	}
	fmt.Println(result)
}

func exercise11() {
	fmt.Println(strings.Join(os.Args, " "))
}

func exercise12() {
	fmt.Print(strings.Join(os.Args, " "))
	for i, arg := range os.Args {
		fmt.Println(string(i) + " " + arg)
	}
}

func exercise13() {
	start := time.Now()
	exercise1()
	timeElapsed := time.Since(start)
	fmt.Printf("time for inefficent %s\n", timeElapsed)

	start = time.Now()
	exercise11()
	timeElapsed = time.Since(start)
	fmt.Printf("time for efficient %s\n", timeElapsed)
}

func exercise14() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.OpenFile(arg, 0, 0777)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, val := range counts {
		if val > 1 {
			fmt.Println(f.Name())
		}
	}
}
