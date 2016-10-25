package controllers

import (
	"net/http"

	"github.com/apiseedprojects/go/responses"
)

type StatusController struct{}

func (sc StatusController) Statusz(w http.ResponseWriter, r *http.Request) {
	responses.OK(w, responses.M{"status": "ok"})
}
