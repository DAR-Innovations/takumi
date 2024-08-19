package routes

import (
	"takumi/internal/modules/authorization"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, handler *authorization.Handler) {
	router := r.Group("/api/v1/auth")

	router.POST("/login", handler.LoginHandler)
	router.POST("/signup", handler.SignUpHandler)
}
