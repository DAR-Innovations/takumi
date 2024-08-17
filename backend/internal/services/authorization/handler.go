package authorization

import (
	"log"
	"net/http"
	"takumi/internal/services/authorization/types"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.Service.Login(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Success!")
}

func (h *Handler) SignUpHandler(c *gin.Context) {
	req := &types.SignupReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Printf("Request body: %v", req) // Log request body
		log.Printf("Bind error: %v", err)   // Log binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.Service.Signup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Success!")
}
