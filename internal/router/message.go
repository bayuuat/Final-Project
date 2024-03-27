package router

import (
	"go-mygram/internal/handler"
	"go-mygram/internal/middleware"

	"github.com/gin-gonic/gin"
)

type MessageRouter interface {
	Mount()
}

type messageRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.MessageHandler
}

func NewMessageRouter(v *gin.RouterGroup, handler handler.MessageHandler) MessageRouter {
	return &messageRouterImpl{v: v, handler: handler}
}

func (m *messageRouterImpl) Mount() {
	m.v.Use(middleware.CheckAuthBearer)

	m.v.POST("/messages", m.handler.CreateMessage)
	m.v.GET("/messages/user", m.handler.GetMessagesByUserID)
	m.v.GET("/messages/photo/:photo_id", m.handler.GetMessagesByPhotoID)
	m.v.PUT("/messages/:id", m.handler.UpdateMessage)
	m.v.DELETE("/messages/:id", m.handler.DeleteMessage)
}
