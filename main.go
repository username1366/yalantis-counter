package main

import (
	"fmt"
	"net/http"
	"log"
)

const (
	SOCKET = ":8080"
)

type Counter struct {
	Iterator chan uint64
}

func (c *Counter) Iterate() {
	var i uint64
	for {
		i++
		c.Iterator <- i
	}
}

func main() {
	c := Counter{
		Iterator: make(chan uint64, 0),
	}
	go c.Iterate()
	http.Handle("/", http.HandlerFunc(c.CountHandler))
	log.Printf("Listen http server %v", SOCKET)
	log.Fatal(http.ListenAndServe(SOCKET, nil))
}

func (c *Counter) CountHandler(w http.ResponseWriter, r *http.Request) {
	i := fmt.Sprintf("%v", <-c.Iterator)
	_, err := w.Write([]byte(i))
	log.Println(i)
	if err != nil {
		log.Println(err)
	}
}
