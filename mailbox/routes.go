package mailbox

import (
	"fmt"
	"net/http"
)

// '/mailbox/{id}'
func RenderMailbox(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprint(w, "Mailbox ", id)
}

// '/mailbox/'
func RenderLanding(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprint(w, "Mailbox ", id)
}
