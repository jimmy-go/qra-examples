// Package auth contains Authorization functions.
package auth

import (
	"log"
	"net/http"

	"github.com/jimmy-go/qra-examples/checklist/sessions"
)

// Handler authentication middleware.
func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := sessions.UserID(w, r)
		if err != nil {
			log.Printf("Handler : qra : cooktoken : err [%s] path [%s]", err, r.RequestURI)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		log.Printf("Handler : token [%v]", token)

		h.ServeHTTP(w, r)
	})
}
