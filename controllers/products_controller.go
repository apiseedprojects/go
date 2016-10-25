package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/apiseedprojects/go/helpers"
	"github.com/apiseedprojects/go/models"
	"github.com/apiseedprojects/go/responses"
	"github.com/gorilla/mux"
)

type ProductsController struct {
	DDB *sql.DB
}

func (hc ProductsController) List(w http.ResponseWriter, r *http.Request) {
	mp := models.NewProductsModel(hc.DDB)
	productsList, err := mp.List()
	if err != nil {
		logrus.
			WithField("message", err.ErrorMessage).
			Error("Error Getting Product List")
		responses.GenericSimplified(w, err.HTTPCode, err.ErrorMessage)
		return
	}
	logrus.
		WithField("products_count", len(productsList)).
		WithField("products", helpers.ToJSON(productsList)).
		Info("Controller Products List")
	responses.OK(w, responses.M{
		"products_count": len(productsList),
		"products":       productsList,
	})
}

func (hc ProductsController) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Reading Posted Product For Creation")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}
	pf := &models.ProductForm{}
	err = json.Unmarshal(b, pf)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Decoding Posted Product For Creation")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}
	pm := models.NewProductsModel(hc.DDB)
	np, id, merr := pm.Create(pf)
	if err != nil {
		logrus.
			WithField("message", merr.ErrorMessage).
			Error("Error Creating Product")
		responses.GenericSimplified(w, merr.HTTPCode, merr.ErrorMessage)
		return
	}
	logrus.
		WithField("inserted_id", helpers.ToJSON(np)).
		WithField("product", helpers.ToJSON(np)).
		Info("Controller Products Create")
	responses.OK(w, responses.M{
		"inserted_id": id,
		"product":     np,
	})
}

func (hc ProductsController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Decoding Requested Product Product ID For Fetch")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}

	mp := models.NewProductsModel(hc.DDB)
	p, merr := mp.Get(id)
	if merr != nil {
		logrus.
			WithField("message", merr.ErrorMessage).
			WithField("id", id).
			Error("Error Fetching Product")
		responses.GenericSimplified(w, merr.HTTPCode, merr.ErrorMessage)
		return
	}

	logrus.
		WithField("requested_id", id).
		WithField("product", helpers.ToJSON(p)).
		Info("Controller Products Read")
	responses.OK(w, responses.M{
		"requested_id": id,
		"product":      p,
	})
}

func (hc ProductsController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Decoding Requested Product Product ID For Update")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Reading Posted Product For Update")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}
	pf := &models.ProductForm{}
	err = json.Unmarshal(b, pf)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Decoding Posted Product For Update")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}
	pm := models.NewProductsModel(hc.DDB)
	up, ra, merr := pm.Update(id, pf)
	if err != nil {
		logrus.
			WithField("message", merr.ErrorMessage).
			Error("Error Updating Product")
		responses.GenericSimplified(w, merr.HTTPCode, merr.ErrorMessage)
		return
	}
	if ra == 0 {
		logrus.
			WithField("message", "no rows affected").
			Error("Error Updating Product")
		responses.GenericSimplified(w, http.StatusNotFound, "no rows affected with update")
		return
	}
	logrus.
		WithField("updated_id", id).
		WithField("rows_affected", ra).
		WithField("product", helpers.ToJSON(up)).
		Info("Controller Products Create")
	responses.OK(w, responses.M{
		"inserted_id":   id,
		"rows_affected": ra,
		"product":       up,
	})
}

func (hc ProductsController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("Error Decoding Requested Product Product ID For Delete")
		responses.GenericSimplified(w, http.StatusBadRequest, err.Error())
		return
	}

	mp := models.NewProductsModel(hc.DDB)
	merr := mp.Delete(id)
	if merr != nil {
		logrus.
			WithField("message", merr.ErrorMessage).
			WithField("id", id).
			Error("Error Deleting Product")
		responses.GenericSimplified(w, merr.HTTPCode, merr.ErrorMessage)
		return
	}

	logrus.
		WithField("deleted_id", id).
		Info("Controller Products Delete")
	responses.OK(w, responses.M{
		"requested_id": id,
	})
}
