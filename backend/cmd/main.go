package main

import (
	"log"
	"takumi/api"
	"takumi/internal/config"
	"takumi/internal/database"
	"takumi/internal/routes"
	"takumi/internal/services/authorization"

	"github.com/gin-gonic/gin"
)

func main() {
	dbHandler := database.InitDB(config.DBUrl)

	app := gin.Default()
	app.Use(api.Config())
	router := routes.NewTakumiRouter(app, "/api", "/v1")

	authService, err := authorization.InitAuthService(dbHandler)
	if err != nil {
		log.Fatal("error initializing authorization service: ", err)
		return
	}
	authHandlers := authorization.NewHandler(authService)
	router.RegisterAuthRoutes(authHandlers)

	err = app.Run(config.PORT)
	if err != nil {
		log.Fatal("error running server: ", err)
	}
}
