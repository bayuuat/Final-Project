package router

import (
	"go-mygram/internal/handler"
	"go-mygram/internal/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

func (u *userRouterImpl) Mount() {
	u.v.POST("/users/register", u.handler.UserSignUp)
	u.v.POST("/users/login", u.handler.UserSignIn)

	u.v.Use(middleware.CheckAuthBearer)

	u.v.GET("/users", u.handler.GetUsers)
	u.v.PUT("/users", u.handler.UpdateUserByID)
	u.v.DELETE("/users", u.handler.DeleteUsersById)
}
