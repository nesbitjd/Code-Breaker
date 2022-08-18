package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nesbitjd/hangle_server/database"
	"github.com/nesbitjd/hangle_server/router"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	level, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		level = "info"
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.DebugLevel
	}
	logrus.SetLevel(lvl)

	port := "8080"

	logrus.Info("Setting up database")
	err = database.Setup()
	if err != nil {
		logrus.Fatalf("unable to setup database: %w", err)
	}

	logrus.Debug("Intializing gin engine")
	r := gin.Default()

	logrus.Trace("Loading router")
	router := router.Load()

	logrus.Trace("Initializing http server")
	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: router}

	logrus.Info("Starting HTTP server...")
	err = srv.ListenAndServe()
	if err != nil {
		logrus.Fatalf("unable to listen and serve: %w", err)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
