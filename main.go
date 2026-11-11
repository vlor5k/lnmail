package main

import (
	"bakery/lnmail/data"
	"bakery/lnmail/mailbox"
	"fmt"
	"net/http"
)

// import sqlite "gosqlite.org"

// db, _ := sqlite.OpenWAL("app.db") // WAL + busy_timeout=5s + foreign_keys=on
// defer db.Close()

func main() {
	data.SetupDB()

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

	http.ListenAndServe(":8080", mux)
}
