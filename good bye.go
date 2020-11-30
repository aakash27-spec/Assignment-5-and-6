package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// https://golang.org/pkg/net/http/#Handler
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Goodbye request")

	// write the response
	fmt.Fprintf(rw, "Goodbye")
}
