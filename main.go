package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

var (
	goVersion = runtime.Version()
	addr      = ":8080"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "'Hello World' from golang:%s in scratch docker image", goVersion)
}

func main() {
	http.HandleFunc("/", helloHandler)

	exitCh := make(chan string)

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			exitCh <- err.Error()
		}
		return
	}()

	log.Printf("Service started with golang:%s, serving at %s", goVersion, addr)

	for {
		select {
		case errMsg := <-exitCh:
			log.Fatal(errMsg)
			return
		}
	}
}
