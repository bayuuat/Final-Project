package router

import (
	"go-mygram/internal/handler"
	"go-mygram/internal/middleware"

	"github.com/gin-gonic/gin"
)

type PhotoRouter interface {
	Mount()
}

type photoRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.PhotoHandler
}

func NewPhotoRouter(v *gin.RouterGroup, handler handler.PhotoHandler) PhotoRouter {
	return &photoRouterImpl{v: v, handler: handler}
}

func (p *photoRouterImpl) Mount() {
	// Authenticated routes
	p.v.Use(middleware.CheckAuthBearer)
	p.v.GET("/photos", p.handler.GetPhotos)
	p.v.GET("/photos/:id", p.handler.GetPhotoByID)
	p.v.POST("/photos", p.handler.CreatePhoto)
	p.v.PUT("/photos/:id", p.handler.UpdatePhoto)
	p.v.DELETE("/photos/:id", p.handler.DeletePhotoByID)
}
