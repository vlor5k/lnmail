package mailbox

import (
	"fmt"
	"html/template"
	"net/http"
)

// '/mailbox/{id}'
func RenderMailbox(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprint(w, "Mailbox ", id)
}

// '/mailbox/'

type Stats struct {
	RandomCount int64
	CustomCount int64
	Count       int64
	Sats        int64
}

func RenderLanding(w http.ResponseWriter, r *http.Request) {
	stats := Stats{45, 5, 50, 50000}
	template, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprint(w, "Failed to parse template")
		return
	}

	template.Execute(w, stats)
	// fmt.Fprint(w, "Mailbox ", id)
}
