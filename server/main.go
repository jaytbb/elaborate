package main

import (
	"log"
	"net/http"
)

func main() {
	// XXX make doc directory be an arg
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
