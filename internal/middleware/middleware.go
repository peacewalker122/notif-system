package middleware

import (
	"fmt"

	"notifsys/internal/config"
	"notifsys/pkg/tracer"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Run(r *gin.Engine) {
	tracer.New()
	cfg := config.Get().APP

	fmt.Println(cfg.Name)

	r.Use(gin.Recovery())
	r.Use(otelgin.Middleware(cfg.Name))
	r.Use(Trace())
}
