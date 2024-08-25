package user

import (
	"io"
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

func (h *Handler) UploadProfilePictureHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	file, _, err := c.Request.FormFile("profile_picture")
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: No file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Could not read file", http.StatusInternalServerError)
		return
	}

	picture, err := h.Service.UploadProfilePicture(c, userID, imageData)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Could not upload profile picture", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessJSON(c, picture)
}

func (h *Handler) GetProfilePictureHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	picture, err := h.Service.GetProfilePicture(c, userID)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusNotFound)
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Writer.Write(picture.ImageData)
}

func (h *Handler) DeleteProfilePictureHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))
	err := h.Service.DeleteProfilePicture(c, userID)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendMessageWithStatus(c, "Profile picture deleted successfully", http.StatusOK)
}
