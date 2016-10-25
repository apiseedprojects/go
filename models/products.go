package models

import (
	"database/sql"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/apiseedprojects/go/helpers"
)

type ProductsModel struct {
	DDB *sql.DB
}

type Product struct {
	ID          int64          `json:"id"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       float64        `json:"price"`
}

type ProductForm struct {
	ID          int64   `json:"id,omitempty"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewProductsModel(db *sql.DB) *ProductsModel {
	return &ProductsModel{
		DDB: db,
	}
}

func (pm *ProductsModel) List() ([]Product, *ModelError) {
	rows, err := pm.DDB.Query("SELECT * FROM products;")
	if err != nil {
		return nil, NewModelError(http.StatusBadGateway, "error querying database: %s", err.Error())
	}

	pl := []Product{}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.ID, &p.Code, &p.Name, &p.Description, &p.Price)
		if err != nil {
			logrus.
				WithField("message", err.Error()).
				Error("DB Row Scan Error")
			return nil, NewModelError(http.StatusConflict, "error scanning rows: %s", err.Error())
		}
		pl = append(pl, p)
	}
	logrus.
		WithField("products", helpers.ToJSON(pl)).
		Info("Model Products List")
	return pl, nil
}

func (pm *ProductsModel) Get(id int64) (*Product, *ModelError) {
	p := &Product{}
	row := pm.DDB.QueryRow("SELECT * FROM products WHERE id = ?;", id)
	err := row.Scan(&p.ID, &p.Code, &p.Name, &p.Description, &p.Price)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Row Fetch or Scan Error")
		return nil, NewModelError(http.StatusNotFound, "error scanning or fetching row: %s", err.Error())
	}
	logrus.
		WithField("product", helpers.ToJSON(p)).
		Info("Model Product")
	return p, nil
}

func (pm *ProductsModel) Create(p *ProductForm) (*ProductForm, int64, *ModelError) {
	res, err := pm.DDB.Exec("INSERT INTO products (code, name, description, price) values (?, ?, ?, ?);",
		p.Code,
		p.Name,
		p.Description,
		p.Price,
	)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Insert Row Error")
		return nil, 0, NewModelError(http.StatusBadRequest, "error inserting row: %s", err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Insert Getting Inserted ID Error")
		return nil, 0, NewModelError(http.StatusBadRequest, "error inserting row while determining inserted id error: %s", err.Error())
	}
	p.ID = id
	logrus.
		WithField("product", helpers.ToJSON(p)).
		WithField("inserted_id", id).
		Info("Row inserted successfully")
	return p, id, nil
}

func (pm *ProductsModel) Update(id int64, p *ProductForm) (*ProductForm, int64, *ModelError) {
	p.ID = id
	res, err := pm.DDB.Exec("UPDATE products SET code=?, name=?, description=?, price=? WHERE id=?;",
		p.Code,
		p.Name,
		p.Description,
		p.Price,
		id,
	)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Update Row Error")
		return nil, 0, NewModelError(http.StatusBadRequest, "error updating row: %s", err.Error())
	}
	ra, err := res.RowsAffected()
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Update Error While Getting Afftected Rows")
		return nil, 0, NewModelError(http.StatusBadRequest, "error updating row while determining affected rows: %s", err.Error())
	}
	logrus.
		WithField("product", helpers.ToJSON(p)).
		WithField("updated_id", id).
		WithField("rows_affected", ra).
		Info("Row updated successfully")
	return p, ra, nil
}

func (pm *ProductsModel) Delete(id int64) *ModelError {
	res, err := pm.DDB.Exec("DELETE FROM products WHERE id = ?;", id)
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Delete Row Error")
		return NewModelError(http.StatusBadRequest, "error deleting row: %s", err.Error())
	}
	ra, err := res.RowsAffected()
	if err != nil {
		logrus.
			WithField("message", err.Error()).
			Error("DB Delete Error While Getting Afftected Rows")
		return NewModelError(http.StatusBadRequest, "error deleting row while determining affected rows: %s", err.Error())
	}
	if ra == 0 {
		logrus.
			WithField("message", "no rows affected").
			Error("DB Delete Error No Rows Affected, Nothing Was Deleted")
		return NewModelError(http.StatusBadRequest, "error deleting row no affected rows: %d", ra)
	}
	logrus.
		WithField("deleted_id", id).
		WithField("rows affected", ra).
		Info("Row deleted successfully")
	return nil
}
