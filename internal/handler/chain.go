package handler

import "net/http"

type Middleware func(next http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for _, mw := range middlewares {
			final = mw(final)
		}
		return final
	}
}
