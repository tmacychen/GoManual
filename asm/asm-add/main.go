package main

import "fmt"

// no function body!
func add(x, y int64) int64

func main() {
	fmt.Printf("use asm function \"add\" to add tow number up : %v + %v = %v\n", 1, 2, add(1, 2))
}
