package main

import (
	"io"
	"net/http"
)

func main() {
	url := "https://jsonip.com"

	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	response.Body.Close()

	responseBody, bodyError := io.ReadAll(response.Body)

	if bodyError != nil {
		panic(bodyError)
	}

	println(string(responseBody))

}
