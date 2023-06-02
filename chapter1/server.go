package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n, counter: %d", r.URL.Path, count)
}

func exercise112(w http.ResponseWriter, r *http.Request) {
	f, _ := strconv.ParseFloat(os.Args[1], 64)
	exercise15(w, f)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "counter: %d", count)
	mu.Unlock()
}
