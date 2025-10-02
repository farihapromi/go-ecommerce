package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // safer than *
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Promi")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down to the actual mux
		mux.ServeHTTP(w, r)
	})
}
