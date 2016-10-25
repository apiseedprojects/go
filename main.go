package main

import (
	"flag"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/apiseedprojects/go/bootstrap"
	"github.com/apiseedprojects/go/dependencies/dbconn"
	"github.com/apiseedprojects/go/middlewares"
	"github.com/urfave/negroni"
)

func main() {
	configFilePath := ""
	useTLS := false

	flag.StringVar(&configFilePath, "config", "local/config.json", "path to config file (default \"local/config.json\")")
	flag.BoolVar(&useTLS, "tls", false, "use TLS (HTTPS) in server (default false)")
	flag.Parse()

	useTLS, config := bootstrap.ReadConfig(configFilePath, useTLS)

	bindAddress := config.BindAddress

	ddb, gdb, err := dbconn.GetDB(config.DBConnectionString)
	if err != nil {
		panic(err)
	}

	routes := bootstrap.Setup(&bootstrap.AppDependencies{
		DDB:           ddb,
		GDB:           gdb,
		JWTSigningKey: config.JWTSigningKey,
	})

	jsonContentTypeMiddleware := middlewares.NewGlobalHeaders("Content-Type", "application/json")

	n := negroni.Classic()
	n.Use(jsonContentTypeMiddleware)
	n.UseHandler(routes)

	logrus.
		WithField("bind_address", bindAddress).
		WithField("use_tls", useTLS).
		Info("Starting Server")

	if useTLS {
		err = http.ListenAndServeTLS(bindAddress, "ssl/apiseedprojectsgo.crt", "ssl/apiseedprojectsgo.key", n)
	} else {
		err = http.ListenAndServe(bindAddress, n)
	}

	if err != nil {
		panic(err)
	}
}
