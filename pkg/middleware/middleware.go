package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if len(bearer) < 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		split := strings.Split(bearer, "Bearer ")
    token := split[1]

		fmt.Println(token)

		HandlerFunc.ServeHTTP(w, r)
	})
}
