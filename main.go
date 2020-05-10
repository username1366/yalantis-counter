package main

import (
	"fmt"
	"net/http"
	"sync"
	"log"
)

const (
	SOCKET = ":8080"
)

type Counter struct {
	mu sync.Mutex
	i uint64
}

func (c *Counter) Get() uint64 {
	c.mu.Lock()
	c.i++
	c.mu.Unlock()
	return c.i
}

func main() {
	c := Counter{}
	http.Handle("/", http.HandlerFunc(c.CountHandler))
	log.Printf("Listen http server %v", SOCKET)
	log.Fatal(http.ListenAndServe(SOCKET, nil))
}

func (c *Counter) CountHandler(w http.ResponseWriter, r *http.Request) {
	i := fmt.Sprintf("%v", c.Get())
	_, err := w.Write([]byte(i))
	log.Println(i)
	if err != nil {
		log.Println(err)
	}
}
