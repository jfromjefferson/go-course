package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	jsonContent := bytes.NewBuffer([]byte(`{"name": "Golang", "description": "Golang course, from zero to hero", "price": 100}`))
	resp, err := client.Post("https://www.google.com", "text/plain", jsonContent)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
