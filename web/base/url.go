package main

import (
	"fmt"
	"io"
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
			url = "https://" + url
		}

		resp, err = http.Get(url)
		fmt.Println(url)
		if err != nil {
			if resp != nil {
				fmt.Fprintf(os.Stderr, "fetch :%v\n Status:%v\n", err, resp.Status)
			} else {
				fmt.Fprintf(os.Stderr, "fetch :%v\n", err)
			}
			os.Exit(1)
		}

		//		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :reading %s :%v\n Status :%v\n", url, err, resp.Status)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "io.Copy to stdout err :%v\n", err)
		}
		fmt.Printf("Status :%v\n", resp.Status)
		resp.Body.Close()
	}
}
