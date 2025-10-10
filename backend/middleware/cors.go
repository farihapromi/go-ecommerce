package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // safer than *
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Promi")

		// Pass down to the actual mux
		next.ServeHTTP(w, r)
	})
}
