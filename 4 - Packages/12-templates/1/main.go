package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name        string
	Description string
	Price       int
}

func main() {
	course := Course{
		Name:        "Golang",
		Description: "Golang course, from zero to hero",
		Price:       100,
	}
	tmp := template.New("courseTemplate")
	tmp, err := tmp.Parse("Course: {{.Name}}\nDescription: {{.Description}}\nPrice: {{.Price}}\n")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
