package main

import (
	"io"
	"log"
	"net/http"
)

var (
	MaxWorker = 10
	MaxQueue  = 1
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	work := Job{Payload: &Payload{123}}

	// Push the work onto the queue.
	JobQueue <- work

	io.WriteString(w, "hello world!")

}

func main() {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}
