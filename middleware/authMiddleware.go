package middleware

import (
	"net/http"
	"time"

	"github.com/ramonaekanayake/grit-space/models"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		var expiresAt time.Time
		err = models.DB.QueryRow("SELECT expires_at FROM sessions WHERE session_token = ?", sessionToken.Value).Scan(&expiresAt)
		if err != nil || expiresAt.Before(time.Now()) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
