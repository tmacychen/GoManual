package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	a := reflect.ValueOf(&x)
	m := a.Elem()
	m.SetFloat(3.1415)
	fmt.Println("value change is ", m)

}
