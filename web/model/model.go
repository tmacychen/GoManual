package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("filedname example")
	t, _ = t.Parse("hello {{.UserName}}! \n")
	p := Person{UserName: "Tmacy"}
	t.Execute(os.Stdout, p)
}
