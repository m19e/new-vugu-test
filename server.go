// +build ignore

package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("dist")))

	if err := http.ListenAndServe(":8686", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
