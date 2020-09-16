package middlewares

import "net/http"

// ContentTypeMiddlewareFunc injects the content type header to let the user know its response is JSON
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
