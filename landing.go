package main

import (
	"fmt"
	"net/http"
)

func landingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi")
}
