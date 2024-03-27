package handler

import (
	"net/http"
	"strconv"

	"go-mygram/internal/model"
	"go-mygram/internal/service"

	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	GetPhotos(ctx *gin.Context)
	GetPhotoByID(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhotoByID(ctx *gin.Context)
	CreatePhoto(ctx *gin.Context)
}

type photoHandlerImpl struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{
		photoService: photoService,
	}
}

func (h *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	photos, err := h.photoService.GetPhotos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, photos)
}

func (h *photoHandlerImpl) GetPhotoByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	photo, err := h.photoService.GetPhotoByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, photo)
}

func (h *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	var updatedPhoto model.PhotoPost
	if err := ctx.BindJSON(&updatedPhoto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPhotoResult, err := h.photoService.UpdatePhoto(ctx, id, updatedPhoto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedPhotoResult)
}

func (h *photoHandlerImpl) DeletePhotoByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	err = h.photoService.DeletePhotoByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

func (h *photoHandlerImpl) CreatePhoto(ctx *gin.Context) {
	var photo model.PhotoPost
	if err := ctx.BindJSON(&photo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil userID dari JWT header
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi userID menjadi uint64
	userIDUint64, ok := userID.(uint64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	createdPhoto, err := h.photoService.CreatePhoto(ctx, userIDUint64, photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdPhoto)
}
