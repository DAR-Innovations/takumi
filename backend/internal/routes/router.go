package routes

import (
	"takumi/internal/modules/authorization"

	"github.com/gin-gonic/gin"
)

type TakumiRouter struct {
	Prefix  string
	Version string
	Engine  *gin.Engine
}

func NewTakumiRouter(engine *gin.Engine, prefix string, version string) *TakumiRouter {
	return &TakumiRouter{
		Prefix:  prefix,
		Version: version,
		Engine:  engine,
	}
}

func (tr *TakumiRouter) RegisterAuthRoutes(handler *authorization.Handler) {
	router := tr.Engine.Group("/auth")

	router.POST("/login", handler.LoginHandler)
	router.POST("/signup", handler.SignUpHandler)
	router.GET("/current-user", handler.GetCurrentUser)
}
