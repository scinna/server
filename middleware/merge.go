package middleware

import (
	"net/http"

	"github.com/scinna/server/services"
)

// CombineMiddlewares combines all middleware, pretty much equivalent to what you would find in a react app
func CombineMiddlewares(prv *services.Provider, handlerfunc http.HandlerFunc, setContentType bool) http.HandlerFunc {

	var hd http.HandlerFunc = handlerfunc

	if setContentType {
		hd = ContentTypeMiddlewareFunc(hd)
	}

	hd = RateLimiterMiddlewareFunc(prv, hd)

	return hd
}

// CombineMiddlewaresCT combines all middleware, pretty much equivalent to what you would find in a react app
func CombineMiddlewaresCT(prv *services.Provider, handlerfunc http.HandlerFunc) http.HandlerFunc {
	return CombineMiddlewares(prv, handlerfunc, true)
}
