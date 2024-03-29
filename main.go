package main

import (
	"fmt"
	"net/http"
	"time"

	"go-mygram/internal/handler"
	"go-mygram/internal/infrastructure"
	"go-mygram/internal/model"
	"go-mygram/internal/repository"
	"go-mygram/internal/router"
	"go-mygram/internal/service"
	"go-mygram/pkg"
	"go-mygram/pkg/helper"

	"github.com/gin-gonic/gin"

	_ "go-mygram/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DUCUMENTATION
// @version		2.0
// @description	golong kominfo 006 api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	server()
}

func server() {
	g := gin.Default()
	g.Use(gin.Recovery())

	// /public => generate JWT public
	g.GET("/public", func(ctx *gin.Context) {
		now := time.Now()

		claim := model.StandardClaim{
			Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
			Iss: "go-middleware",
			Aud: "golang-006",
			Sub: "public-token",
			Exp: uint64(now.Add(time.Hour).Unix()),
			Iat: uint64(now.Unix()),
			Nbf: uint64(now.Unix()),
		}
		token, err := helper.GenerateToken(claim)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
				Message: "error generating public token",
				Errors:  []string{err.Error()},
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]any{"token": token})
	})

	usersGroup := g.Group("/api")

	gorm := infrastructure.NewGormPostgres()
	userRepo := repository.NewUserQuery(gorm)
	userSvc := service.NewUserService(userRepo)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	userRouter.Mount()

	photoRepo := repository.NewPhotoRepository(gorm)
	photoSvc := service.NewPhotoService(photoRepo)
	photoHdl := handler.NewPhotoHandler(photoSvc)
	photoRouter := router.NewPhotoRouter(usersGroup, photoHdl)

	photoRouter.Mount()

	messageRepo := repository.NewMessageRepository(gorm)
	messageSvc := service.NewMessageService(messageRepo)
	messageHdl := handler.NewMessageHandler(messageSvc)
	messageRouter := router.NewMessageRouter(usersGroup, messageHdl)

	messageRouter.Mount()

	sosmedRepo := repository.NewSocialMediaRepository(gorm)
	sosmedSvc := service.NewSocialMediaService(sosmedRepo)
	sosmedHdl := handler.NewSocialMediaHandler(sosmedSvc)
	sosmedRouter := router.NewSocialMediaRouter(usersGroup, sosmedHdl)

	sosmedRouter.Mount()

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}
