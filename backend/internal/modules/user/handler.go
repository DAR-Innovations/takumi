package user

import (
	"net/http"
	"strconv"
	"takumi/internal/modules/user/types"
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

func (h *Handler) GetUserByIDHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUserByID(c, userID)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusNotFound)
		return
	}

	utils.SendSuccessJSON(c, user)
}

func (h *Handler) DeleteUserByIDHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUserByID(c, userID)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusNotFound)
		return
	}

	utils.SendMessageWithStatus(c, "User deleted successfully", http.StatusOK)
}

func (h *Handler) UpdateUserParamsHandler(c *gin.Context) {
	update := &types.User{}
	if err := c.ShouldBindJSON(update); err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.Service.UpdateUserParams(c, *update)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessJSON(c, updatedUser)
}
