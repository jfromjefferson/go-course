package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8000...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	mux.Handle("/blog", Blog{})
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Blog struct{}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Blog"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
