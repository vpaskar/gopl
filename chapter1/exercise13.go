package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get("http://" + url)
		if err != nil {
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		if err := resp.Body.Close(); err != nil {
			os.Exit(1)
		}
		fmt.Print(string(b))
	}
}

func exercise17(out io.Writer) {
	for _, url := range os.Args[1:] {
		resp, err := http.Get("http://" + url)
		if err != nil {
			os.Exit(1)
		}
		_, err = io.Copy(out, resp.Body)
		if err := resp.Body.Close(); err != nil {
			os.Exit(1)
		}
	}
}

func exercise18(out io.Writer) {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		_, err = io.Copy(out, resp.Body)
		if err := resp.Body.Close(); err != nil {
			os.Exit(1)
		}
	}
}

func exercise19(out io.Writer) {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		_, err = io.Copy(out, resp.Body)
		out.Write([]byte("\nStatusCode: " + string(resp.Status)))
		if err := resp.Body.Close(); err != nil {
			os.Exit(1)
		}
	}
}
