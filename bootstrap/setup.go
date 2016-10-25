package bootstrap

import (
	"database/sql"
	"net/http"

	"github.com/apiseedprojects/go/controllers"
	"github.com/apiseedprojects/go/middlewares"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

type AppDependencies struct {
	DDB           *sql.DB
	GDB           *gorm.DB
	JWTSigningKey string
}

func Setup(deps *AppDependencies) http.Handler {
	router := mux.NewRouter()

	homeController := &controllers.HomeController{}
	statusController := &controllers.StatusController{}
	authController := &controllers.AuthController{JWTSigningKey: deps.JWTSigningKey}
	productsController := &controllers.ProductsController{DDB: deps.DDB}

	router.HandleFunc("/", homeController.Home).Methods("GET")
	router.HandleFunc("/statusz", statusController.Statusz).Methods("GET")
	router.HandleFunc("/auth", authController.Login).Methods("GET")

	jwtAuthMiddleware := middlewares.NewJWTAuth(deps.JWTSigningKey)

	v1router := mux.NewRouter().PathPrefix("/v1").Subrouter().StrictSlash(true)

	v1router.HandleFunc("/products", productsController.List).Methods("GET")
	v1router.HandleFunc("/products", productsController.Create).Methods("POST")
	v1router.HandleFunc("/products/{id:[0-9]+}", productsController.Read).Methods("GET")
	v1router.HandleFunc("/products/{id:[0-9]+}", productsController.Update).Methods("PUT")
	v1router.HandleFunc("/products/{id:[0-9]+}", productsController.Delete).Methods("DELETE")

	router.
		PathPrefix("/v1").
		Handler(
			negroni.New(
				// Other middleware here.
				jwtAuthMiddleware,
				negroni.Wrap(v1router),
			),
		)

	return router
}
