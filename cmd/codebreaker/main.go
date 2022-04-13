package main

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	port := "8080"

	err := database.Setup()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	router := router.Load()

	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: router}

	logrus.Info("Starting HTTP server...")
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
