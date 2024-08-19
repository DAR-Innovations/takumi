package authorization

import (
	"net/http"
	"takumi/internal/middleware"
	"takumi/internal/modules/authorization/types"
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

	middleware.SetCookieHandler(c, currentUser.Token)
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

	middleware.SetCookieHandler(c, currentUser.Token)
	utils.SendSuccessJSON(c, currentUser)
}
