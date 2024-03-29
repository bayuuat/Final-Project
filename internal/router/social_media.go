package router

import (
	"go-mygram/internal/handler"
	"go-mygram/internal/middleware"

	"github.com/gin-gonic/gin"
)

type SocialMediaRouter interface {
	Mount()
}

type socialMediaRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.SocialMediaHandler
}

func NewSocialMediaRouter(v *gin.RouterGroup, handler handler.SocialMediaHandler) SocialMediaRouter {
	return &socialMediaRouterImpl{
		v:       v,
		handler: handler,
	}
}

func (r *socialMediaRouterImpl) Mount() {
	socialMediaGroup := r.v.Group("/socialmedias")
	r.v.Use(middleware.CheckAuthBearer)
	{
		socialMediaGroup.POST("/", r.handler.CreateSocialMedia)
		socialMediaGroup.GET("/:id", r.handler.GetSocialMediaByID)
		socialMediaGroup.PUT("/:id", r.handler.UpdateSocialMedia)
		socialMediaGroup.DELETE("/:id", r.handler.DeleteSocialMediaByID)
	}
}
