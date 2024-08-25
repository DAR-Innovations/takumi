package routes

import (
	"takumi/internal/modules/authorization"
	"takumi/internal/modules/user"

	"github.com/gin-gonic/gin"
)

type TakumiRouter struct {
	Prefix  string
	Version string
	Routes  *gin.RouterGroup
}

func NewTakumiRouter(engine *gin.Engine, prefix string, version string) *TakumiRouter {
	router := engine.Group(prefix + version)
	return &TakumiRouter{
		Prefix:  prefix,
		Version: version,
		Routes:  router,
	}
}

func (tr *TakumiRouter) RegisterAuthRoutes(handler *authorization.Handler) {
	router := tr.Routes.Group("/auth")

	router.POST("/login", handler.LoginHandler)
	router.POST("/signup", handler.SignUpHandler)
	router.GET("/current-user", handler.GetCurrentUser)
}

func (tr *TakumiRouter) RegisterUserRoutes(handler *user.Handler) {
	router := tr.Routes.Group("/user")

	router.GET("/:id", handler.GetUserByIDHandler)
	router.DELETE("/:id", handler.DeleteUserByIDHandler)
	router.PUT("/update", handler.UpdateUserParamsHandler)
	router.POST("/:id/profile-picture", handler.UpdateProfilePictureHandler)
	router.GET("/:id/profile-picture", handler.GetProfilePictureByUserID)
	router.DELETE("/:id/profile-picture", handler.DeleteProfilePictureHandler)
}
