package server

import (
	"net/http"

	"notifsys/internal/app/notif"
	"notifsys/internal/app/user"
	"notifsys/internal/factory"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"notifsys/docs"
)

func Run(r *gin.Engine, f *factory.Factory) {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = "localhost:300"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), gin.BasicAuth(gin.Accounts{
		"admin": "password",
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		user.NewHandler(f).Route(v1.Group("/user"))
		notif.NewHandler(f).Route(v1.Group("/notif"))
	}
}
