package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
}

func main() {
	jhon := Student{Human{"Jhon", 33, 190}, "Computer Science"}

	fmt.Println("His name is ", jhon.name)
	fmt.Println("His age  is ", jhon.age)
	fmt.Println("His weight is ", jhon.weight)

	jhon.speciality = "AI"
	fmt.Println("Jhon changed his speciality")
	fmt.Println("Jhon 's speciality is ", jhon.speciality)
	jhon.age = 66
	fmt.Println("jhon changed his age ,his age is ", jhon.age)
}
