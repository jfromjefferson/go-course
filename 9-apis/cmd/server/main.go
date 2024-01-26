package main

import (
	"fmt"
	"github.com/jfromjefferson/gi-course-9/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.DBName)
}
