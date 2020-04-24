// +build ignore

package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("dist")))

	log.Printf("Starting HTTP Server at %q", "127.0.0.1:8686")

	if err := http.ListenAndServe(":8686", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
