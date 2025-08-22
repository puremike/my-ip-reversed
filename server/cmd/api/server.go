package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/puremike/ip_reversed/util"
)

type healthChecker struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (app *application) route() http.Handler {

	hC := &healthChecker{
		Status:  "OK",
		Message: "Service is running",
	}

	g := gin.Default()

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{util.GetEnvString("FRONTEND_URL", "http://localhost:3030")},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := g.Group("/v1")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, hC)
	})

	v1.GET("/myip", app.ipReversedHandler)
	return g
}

func (app *application) server(mux http.Handler) error {

	server := &http.Server{
		Addr:    ":" + app.AppConfig.PORT,
		Handler: mux,
	}

	log.Printf("Starting server on port %s", app.AppConfig.PORT)

	return server.ListenAndServe()
}
