package main

import (
	"bakery/lnmail/mailbox"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, world")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "hi")
	// })

	mux := http.NewServeMux()

	// mux.HandleFunc("GET /{$}", landingHandler)
	mux.HandleFunc("GET /mailbox/{id}", mailbox.RenderMailbox)
	mux.HandleFunc("GET /{$}", mailbox.RenderLanding)

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", mux)
}
