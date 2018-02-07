package main

import "fmt"

// no function body!
func rotl(x, y int64) int64

func main() {
	var i int64
	i = 1
	fmt.Printf("use asm function \"rol\" to rotate left : %v << %v = %v\n", 2, i, rotl(2, i))
}
