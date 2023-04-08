package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}
