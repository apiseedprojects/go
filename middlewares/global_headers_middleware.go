package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

type GlobalHeaders struct{}

func NewGlobalHeaders(key string, value string) negroni.HandlerFunc {
	return negroni.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			w.Header().Set(key, value)
			next(w, r)
		},
	)
}
