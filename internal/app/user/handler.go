package user

import (
	"net/http"

	"notifsys/internal/dto"
	"notifsys/internal/factory"
	"notifsys/internal/service"
	"notifsys/internal/service/interfaces"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service interfaces.User
}

func NewHandler(f *factory.Factory) *handler {
	service.NewUser(f)
	return &handler{
		service: service.NewUser(f),
	}
}

// @Summary      CreateUser
// @Description  CreateUser
// @Tags         User
// @Accept       json
// @Produce      json
// @Param request body dto.User true "request body"
// @Success      200  {object}  dto.SignupRequest
// @Failure      400  {object}  string "Bad Request"
// @Failure      404  {object}  string "Not Found"
// @Failure      500  {object}  string "Internal Server Error"
// @Router       /api/v1/user [post]
func (h *handler) Create(c *gin.Context) {
	payload := new(dto.SignupRequest)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.Create(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &dto.User{
		ID:    data.ID,
		Name:  data.Username,
		Email: data.Email,
	})
}
