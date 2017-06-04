package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	var buffer bytes.Buffer

	s := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString("test is here\n")
	}

	sa := buffer.String()
	e := time.Now()
	fmt.Println("buffer taked time is ", e.Sub(s).Seconds())

	s = time.Now()
	var sl []string
	for i := 0; i < 100000; i++ {
		sl = append(sl, "test is here\n")
	}
	sc := strings.Join(sl, "")
	e = time.Now()
	fmt.Println("append taked time is", e.Sub(s).Seconds())
	if sa == sc {
		fmt.Println("is same")
	}

	sb := ""
	s = time.Now()
	for i := 0; i < 100000; i++ {
		sb += "test is here\n"
	}
	e = time.Now()
	fmt.Println("+= taked time is ", e.Sub(s).Seconds())
	if sb == sa {
		fmt.Println("is same")
	}
}
