package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", GetPostalCode)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe error: %s\n", err)
		return
	}

}

func GetPostalCode(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write error: %s\n", err)
		return
	}
}
