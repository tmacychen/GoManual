package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

//golang 原生log相关

func main() {
	fmt.Println("log test")
	logname := "ll.org"
	logfile, err := os.Create(logname)
	defer logfile.Close()

	if err != nil {
		log.Fatalln("open file now")
	}

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	fmt.Print(&buf)
}
