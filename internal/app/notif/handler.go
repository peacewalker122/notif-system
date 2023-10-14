package notif

import (
	"net/http"

	"notifsys/internal/dto"
	"notifsys/internal/factory"
	"notifsys/internal/service"
	"notifsys/internal/service/interfaces"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service interfaces.Notif
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: service.NewNotif(f),
	}
}

// @Summary      CreateNotif
// @Description  CreateNotif
// @Tags         Notif
// @Accept       json
// @Produce      json
// @Param request body dto.NotifRequest true "request body"
// @Success      200  {object}  string "success"
// @Failure      400  {object}  string "Bad Request"
// @Failure      404  {object}  string "Not Found"
// @Failure      500  {object}  string "Internal Server Error"
// @Router       /api/v1/notif [post]
func (h *handler) Create(c *gin.Context) {
	payload := new(dto.NotifRequest)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Create(c, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
