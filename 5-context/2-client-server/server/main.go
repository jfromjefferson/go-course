package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request finished")
	select {
	case <-time.After(5 * time.Second):
		message := "Request processed successfully"
		log.Println(message)
		w.Write([]byte(message))
	case <-ctx.Done():
		log.Println("Request cancelled by user")
	}

}
