package middlewares

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/urfave/negroni"
)

func NewJWTAuth(signingKey string) negroni.HandlerFunc {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return negroni.HandlerFunc(jwtMiddleware.HandlerWithNext)
}
