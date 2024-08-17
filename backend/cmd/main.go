package main

import (
	"log"
	"takumi/internal/config"
	"takumi/internal/database"
	"takumi/internal/routes"
	"takumi/internal/services/authorization"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbHandler := database.InitDB(config.DBUrl)

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	authService, err := authorization.InitAuthService(dbHandler)
	if err != nil {
		log.Fatal("error initializing authorization service: ", err)
		return
	}
	authHandlers := authorization.NewHandler(authService)
	routes.RegisterAuthRoutes(app, authHandlers)

	err = app.Run(config.PORT)
	if err != nil {
		log.Fatal("error running server: ", err)
	}
}
