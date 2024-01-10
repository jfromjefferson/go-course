package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8000...")
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Blog"))
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
