package user

import (
	"net/http"

	"notifsys/internal/dto"
	"notifsys/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type handler struct {
	service service.User
}

func NewHandler(DB *bun.DB) *handler {
	service.NewUser(DB)
	return &handler{
		service: service.UserService,
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

	data, err := h.service.Create(c, payload)
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
