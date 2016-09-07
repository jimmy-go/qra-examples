// Package auth contains Authorization functions.
package auth

import (
	"log"
	"net/http"

	"github.com/jimmy-go/qra"
)

// MW authentication middleware.
func MW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		log.Printf("Auth : MW : token [%v]", token)

		sess, err := qra.SessionLocate(token)
		if err != nil {
			log.Printf("AuthMW : qra : session : err [%s]", err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		log.Printf("AuthMW : qra : session [%v]", sess)

		h.ServeHTTP(w, r)
	})
}
