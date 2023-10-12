package notif

import "github.com/gin-gonic/gin"

func (r *handler) Route(g *gin.RouterGroup) {
	g.POST("", r.Create)
}
