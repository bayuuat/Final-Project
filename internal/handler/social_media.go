package handler

import (
	"net/http"
	"strconv"

	"go-mygram/internal/model"
	"go-mygram/internal/service"

	"github.com/gin-gonic/gin"
)

type SocialMediaHandler interface {
	CreateSocialMedia(c *gin.Context)
	GetSocialMediaByID(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMediaByID(c *gin.Context)
}

type socialMediaHandlerImpl struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandlerImpl{
		socialMediaService: socialMediaService,
	}
}

func (h *socialMediaHandlerImpl) CreateSocialMedia(c *gin.Context) {
	var socialMediaPost model.SocialMediaPost
	if err := c.ShouldBindJSON(&socialMediaPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil userID dari JWT header
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID menjadi uint64
	userIDUint64, ok := userID.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	socialMedia, err := h.socialMediaService.CreateSocialMedia(c.Request.Context(), userIDUint64, socialMediaPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create social media"})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}

func (h *socialMediaHandlerImpl) GetSocialMediaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	socialMedia, err := h.socialMediaService.GetSocialMediaByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media not found"})
		return
	}

	c.JSON(http.StatusOK, socialMedia)
}

func (h *socialMediaHandlerImpl) UpdateSocialMedia(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	var socialMediaPost model.SocialMediaPost
	if err := c.ShouldBindJSON(&socialMediaPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSocialMedia, err := h.socialMediaService.UpdateSocialMedia(c.Request.Context(), id, socialMediaPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update social media"})
		return
	}

	c.JSON(http.StatusOK, updatedSocialMedia)
}

func (h *socialMediaHandlerImpl) DeleteSocialMediaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media ID"})
		return
	}

	err = h.socialMediaService.DeleteSocialMediaByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete social media"})
		return
	}

	c.JSON(http.StatusNoContent, "You social media has been successfully deleted")
}
