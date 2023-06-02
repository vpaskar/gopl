package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	gifHandler := func(w http.ResponseWriter, r *http.Request) {
		f, _ := strconv.ParseFloat(os.Args[1], 64)
		exercise15(w, f)
	}
	http.HandleFunc("/", gifHandler)
	//http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8082", nil))
}
