package mailbox

import (
	"bakery/lnmail/config"
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

type LandingData struct {
	Stats  Stats
	Config *config.ConfigType
}

func RenderLanding(w http.ResponseWriter, r *http.Request) {
	// var a []int64
	// a = append(a, 4, 66, 3, 1, 9)
	// stats := Stats{45, 5, 50, 50000}
	data := LandingData{
		Stats{45, 5, 50, 50000},
		&config.Config,
	}
	template, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprint(w, "Failed to parse template")
		return
	}

	template.Execute(w, data)
	// fmt.Fprint(w, "Mailbox ", id)
}
