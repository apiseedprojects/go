package controllers

import (
	"net/http"

	"github.com/apiseedprojects/go/responses"
)

type HomeController struct{}

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request) {
	responses.OK(w, responses.M{"app": "apiseedproject"})
}
