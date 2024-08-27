package user

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
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

func (h *Handler) PatchUserParamsHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid user ID", http.StatusBadRequest)
		return
	}

	updateData := map[string]interface{}{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.Service.PatchUserParams(c, userID, updateData)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessJSON(c, updatedUser)
}

func (h *Handler) UpdateProfilePictureHandler(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid user ID", http.StatusBadRequest)
		return
	}

	file, _, err := c.Request.FormFile("profilePicture")
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: File upload failed", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := fmt.Sprintf("user_%d_profile_pic.jpg", userID)
	if err := utils.SaveFile(file, "profile-pictures", fileName); err != nil {
		log.Printf("Failed to save profile picture for user %d: %v", userID, err)
		utils.SendMessageWithStatus(c, "ERROR: Could not save profile picture", http.StatusInternalServerError)
		return
	}

	profilePictureURL := filepath.Join("profile-pictures", fileName)
	updatedUser, err := h.Service.UpdateProfilePicture(c, userID, profilePictureURL)
	if err != nil {
		log.Printf("Error updating profile picture for user %d: %v", userID, err)
		utils.SendMessageWithStatus(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessJSON(c, updatedUser)
}

func (h *Handler) GetProfilePictureByUserID(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid user ID", http.StatusBadRequest)
		return
	}

	profilePictureURL, err := h.Service.GetProfilePictureByUserID(userID)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessJSON(c, gin.H{"profilePictureURL": profilePictureURL})
}

func (h *Handler) DeleteProfilePictureHandler(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.SendMessageWithStatus(c, "ERROR: Invalid user ID", http.StatusBadRequest)
		return
	}

	deletedUser, err := h.Service.DeleteProfilePicture(c, userID)

	fileName := fmt.Sprintf("user_%d_profile_pic.jpg", userID)
	if err := utils.DeleteFile("profile-pictures", fileName); err != nil {
		log.Printf("Failed to delete: %s, error: %v", fileName, err)
		return
	}

	if err != nil {
		utils.SendMessageWithStatus(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessJSON(c, deletedUser)
}
