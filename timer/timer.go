package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	t, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("err :%v \n", err)
	}
	cmd := os.Args[2]

	argstring := strings.Join(os.Args[3:], "")

	sign := timer1(t)
	if ret := <-sign; ret == 1 {
		c := exec.Command(cmd, argstring)
		out, err := c.Output()
		if err != nil {
			log.Fatalf("error :%v\n", err)
		}
		fmt.Println(string(out))

	}
}

func timer_sec() *time.Ticker {
	return time.NewTicker(1 * time.Second)
}
func timer_min() *time.Ticker {
	return time.NewTicker(1 * time.Minute)
}

func timer1(num int64) chan byte {
	sign := make(chan byte, 2)
	t1 := timer_sec()
	go func() {
		for i := num; i > 0; i-- {
			select {
			case <-t1.C:
				//fmt.Printf("time :%v\n", time.Now().Format(time.RFC1123))
			}
		}
		sign <- 1
	}()
	return sign
}
