package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type ServerHandler struct {
}

func (s *ServerHandler) ServeHTTP() {

}

func main() {

	receiver := make(chan interface{})

	go func() {
		for {
			fmt.Println(<-receiver)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		receiver <- html.EscapeString(r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":9999", nil))

	fmt.Println("done")
}
