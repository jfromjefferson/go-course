package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var counter uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		counter++
		m.Unlock()

		w.Write([]byte(fmt.Sprintf("Counter: %d", counter)))
		time.Sleep(300 * time.Millisecond)
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
