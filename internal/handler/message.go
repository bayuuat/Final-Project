package handler

import (
	"net/http"
	"strconv"

	"go-mygram/internal/model"
	"go-mygram/internal/service"

	"github.com/gin-gonic/gin"
)

type MessageHandler interface {
	CreateMessage(ctx *gin.Context)
	GetMessagesByUserID(ctx *gin.Context)
	GetMessagesByPhotoID(ctx *gin.Context)
	UpdateMessage(ctx *gin.Context)
	DeleteMessage(ctx *gin.Context)
}

type messageHandlerImpl struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) MessageHandler {
	return &messageHandlerImpl{
		messageService: messageService,
	}
}

func (h *messageHandlerImpl) CreateMessage(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")

	var message model.MessagePost
	if err := ctx.BindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDUint64, _ := userID.(uint64)
	createdMessage, err := h.messageService.CreateMessage(ctx, userIDUint64, message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdMessage)
}

func (h *messageHandlerImpl) GetMessagesByUserID(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	userIDUint64, _ := userID.(uint64)

	messages, err := h.messageService.GetMessagesByUserID(ctx, userIDUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

func (h *messageHandlerImpl) GetMessagesByPhotoID(ctx *gin.Context) {
	photoIDStr := ctx.Param("photo_id")
	photoID, err := strconv.ParseUint(photoIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	messages, err := h.messageService.GetMessagesByPhotoID(ctx, photoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

func (h *messageHandlerImpl) UpdateMessage(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	var updatedMessage model.MessagePost
	if err := ctx.BindJSON(&updatedMessage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMessageResult, err := h.messageService.UpdateMessage(ctx, id, updatedMessage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedMessageResult)
}

func (h *messageHandlerImpl) DeleteMessage(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message ID"})
		return
	}

	err = h.messageService.DeleteMessage(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}
