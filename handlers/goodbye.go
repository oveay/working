package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye world")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Error in Goodbye handler", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Bye Bye, %s\n", d)
}
