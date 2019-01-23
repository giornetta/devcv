package server

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/giornetta/devcv/auth"
)

func authMiddleware(svc auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")
			username := chi.URLParam(r, "username")

			if err := svc.Authenticate(bearer, username); err != nil {
				respond(w, http.StatusForbidden, e(err.Error()))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
