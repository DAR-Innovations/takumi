package authorization

import (
	"net/http"
	"takumi/internal/config"
	"takumi/internal/services/authorization/types"
	"takumi/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) LoginHandler(c *gin.Context) {
	req := &types.LoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid request body", http.StatusBadRequest)
		return
	}

	currentUser, err := h.Service.Login(c, req)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR:"+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SetCookieHandler(c, utils.CookieSettings{
		Name:     "token",
		Data:     currentUser.Token,
		MaxAge:   3600 * 12,
		Path:     "/",
		Host:     config.HOST,
		Secure:   true,
		HttpOnly: true,
	})
	utils.SendSuccessJSON(c, currentUser)
}

func (h *Handler) SignUpHandler(c *gin.Context) {
	req := &types.SignupReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid request body", http.StatusBadRequest)
		return
	}

	currentUser, err := h.Service.Signup(c, req)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR:"+" "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SetCookieHandler(c, utils.CookieSettings{
		Name:     "token",
		Data:     currentUser.Token,
		MaxAge:   3600 * 12,
		Path:     "/",
		Host:     config.HOST,
		Secure:   true,
		HttpOnly: true,
	})
	utils.SendSuccessJSON(c, currentUser)
}

func (h *Handler) GetCurrentUser(c *gin.Context) {
	token, err := utils.GetCookieHandler(c, utils.CookieSettings{Name: "token"})
	if err != nil {
		utils.SendMessageWithStatus(c, err.Error(), http.StatusNotFound)
		return
	}

	claims, err := utils.ParseToken(*token, config.JWTSecretKey)
	if err != nil {
		utils.SendMessageWithStatus(c, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendSuccessJSON(c, claims)
}
