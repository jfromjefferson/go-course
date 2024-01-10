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
	tmp := template.Must(template.New("courseTemplate").Parse("Course: {{.Name}}\nDescription: {{.Description}}\nPrice: {{.Price}}\n"))
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
