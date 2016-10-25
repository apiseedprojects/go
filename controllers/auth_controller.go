package controllers

import (
	"net/http"
	"time"

	"github.com/apiseedprojects/go/responses"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthController struct {
	JWTSigningKey string
}

func (ac AuthController) Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		responses.Unauthorized(w, responses.M{"message": "invalid or missing authentication user and password"})
		return
	}

	if username != "someuser" || password != "somepass" {
		responses.Unauthorized(w, responses.M{"message": "invalid user/password combination"})
		return
	}

	t := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "apiseedproject",
		"sub": username,
		"exp": t + (24 * 60 * 60),
		"nbf": t,
	})

	tokenString, err := token.SignedString([]byte(ac.JWTSigningKey))
	if err != nil {
		responses.InternalServerError(w, responses.M{"message": err.Error()})
	}

	responses.OK(w, responses.M{"token": tokenString})
}
