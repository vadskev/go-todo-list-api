package auth

import (
	"net/http"

	"github.com/vadskev/go-todo-list-api/internal/config"
)

func New(pass config.PasswordConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(pass.GetPass()) > 0 {
				cookie, err := r.Cookie("token")
				if err != nil {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}

				hash := pass.CreateHash(pass.GetPass())

				if cookie.Value != hash {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
