package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchAll() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go exercise110_2(ch, "http://"+url, "test")
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetchRoutine(ch chan string, url string) {
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprintf("cannot fetch %s", url)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("cannot copy %s", url)
		return
	}
	err = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("cannot close %s", url)
		return
	}
	ch <- fmt.Sprintf("fetching %s, received %d bytes within %v seconds", url, nbytes, elapsed)
}

func exercise110_1() {
	ch := make(chan string)
	for {
		go fetchRoutine(ch, "http://"+os.Args[1])
		fmt.Println(<-ch)
		time.Sleep(2)
	}
}

func exercise110_2(ch chan string, url string, filename string) {
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprintf("cannot fetch %s", url)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	err = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("cannot close %s", url)
		return
	}
	result := fmt.Sprintf("fetching %s, received %d bytes within %v seconds", url, nbytes, elapsed)
	b, err := io.ReadAll(resp.Body)
	err = os.WriteFile(filename, b, 0777)
	if err != nil {
		ch <- fmt.Sprintf("cannot write to file %s", url)
		return
	}
	ch <- result
}
