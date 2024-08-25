package main

import (
	"log"
	"takumi/api"
	"takumi/internal/config"
	"takumi/internal/database"
	"takumi/internal/modules/authorization"
	"takumi/internal/modules/user"
	"takumi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	dbHandler := database.InitDB(cfg.DBSource)

	app := setupGin(cfg)
	app.Static("/uploads", "./uploads")
	router := routes.NewTakumiRouter(app, "/api", "/v1")

	authService, err := authorization.InitAuthService(dbHandler)
	if err != nil {
		log.Fatal("error initializing authorization service: ", err)
		return
	}
	authHandlers := authorization.NewHandler(authService)
	router.RegisterAuthRoutes(authHandlers)

	userService, err := user.InitUserService(dbHandler)
	if err != nil {
		log.Fatal("error initializing user service: ", err)
		return
	}
	userHandlers := user.NewHandler(userService)
	router.RegisterUserRoutes(userHandlers)

	err = app.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("error running server: ", err)
	}
}

func setupGin(cfg *config.Config) *gin.Engine {
	if cfg.Env == "dev" || cfg.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()
	app.Use(api.Config())

	if gin.Mode() == gin.DebugMode {
		app.Use(gin.Logger())
	}

	app.Use(gin.Recovery())

	return app
}
