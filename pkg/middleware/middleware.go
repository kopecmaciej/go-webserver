package middleware

import "net/http"

type authMiddleware struct {
	token string
}

func (auth *authMiddleware) findToken() {
}

func (auth *authMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["token"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
