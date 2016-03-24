package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(`^https?://`)
	var resp *http.Response
	var err error
	for _, url := range os.Args[1:] {
		if !pattern.MatchString(url) {
			url = "http://" + url
		}

		resp, err = http.Get(url)
		fmt.Println(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :%v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :reading %s :%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
