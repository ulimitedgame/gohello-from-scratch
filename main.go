package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go in scratch Docker image")
}

func main() {
	http.HandleFunc("/", helloHandler)

	exitCh := make(chan string)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			exitCh <- err.Error()
		}
		return
	}()

	log.Print("Started, serving at :8080")

	for {
		select {
		case errMsg := <-exitCh:
			log.Fatal(errMsg)
			return
		}
	}
}
