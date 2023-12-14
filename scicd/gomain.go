package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type ServerHandler struct {
}

func (s *ServerHandler) ServeHTTP() {

}

type Listener struct {
}

func (l *Listener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(" in ServeHTTP")
}

func main() {

	receiver := make(chan interface{})
	defer close(receiver)

	go func() {
		for v := range receiver {
			fmt.Println("receive: ", v)
		}
	}()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		receiver <- html.EscapeString(r.URL.Path)
	})

	// log.Fatal(http.ListenAndServe(":9999", nil))

	s := &http.Server{
		Addr:           ":9999",
		Handler:        &Listener{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

	fmt.Println("done")
}
