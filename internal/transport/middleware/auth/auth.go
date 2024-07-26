package auth

import (
	"net/http"

	"github.com/vadskev/go_final_project/internal/lib/logger"
	"github.com/vadskev/go_final_project/internal/lib/utils"
)

func New(pass string) func(next http.Handler) http.Handler {
	const op = "middleware.auth.New"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(pass) > 0 {
				cookie, err := r.Cookie("token")
				if err != nil {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}
				hash := utils.CreateHash(pass)

				if cookie.Value != hash {
					http.Error(w, "Authentification required", http.StatusUnauthorized)
					return
				}
			}
			logger.Info(op)
			next.ServeHTTP(w, r)
		})
	}
}
