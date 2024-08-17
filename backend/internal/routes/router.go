package routes

import (
	"takumi/internal/services/authorization"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, handler *authorization.Handler) {
	router := r.Group("/auth")

	router.POST("/login", handler.LoginHandler)
	router.POST("/signup", handler.SignUpHandler)
}
