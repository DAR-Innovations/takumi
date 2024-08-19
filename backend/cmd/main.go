package main

import (
	"log"
	"takumi/internal/config"
	"takumi/internal/database"
	"takumi/internal/routes"
	"takumi/internal/services/authorization"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbHandler := database.InitDB(cfg.DBSource)

	app := setupGin(cfg)

	authService, err := authorization.InitAuthService(dbHandler)
	if err != nil {
		log.Fatalf("failed to initialize authorization service: %v", err)
	}

	authHandlers := authorization.NewHandler(authService)
	routes.RegisterAuthRoutes(app, authHandlers)

	err = app.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func setupGin(cfg *config.Config) *gin.Engine {
	if cfg.Env == "dev" || cfg.Env == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if gin.Mode() == gin.DebugMode {
		app.Use(gin.Logger())
	}

	app.Use(gin.Recovery())

	return app
}
