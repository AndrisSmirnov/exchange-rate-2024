package http

import (
	user_dto "exchange_rate/pkg/usecase/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTPController) EmailController(router *gin.Engine) {

	emailGroup := router.Group("/")

	emailGroup.Use()
	{
		subGroup := router.Group("/subscribe")
		subGroup.POST("/",
			func(c *gin.Context) {
				ctx := c.Request.Context()

				emailDTO := user_dto.UserEmail{}
				if err := c.ShouldBindJSON(&emailDTO); err != nil {
					c.JSON(http.StatusBadRequest, err.Error())
					return
				}

				if err := h.services.CreateUser(ctx, &user_dto.CreateUserDto{
					UserEmail: emailDTO,
				}); err != nil {
					c.JSON(err.GetCode(), nil)
					return
				}

				c.JSON(http.StatusOK, nil)
			},
		)
	}
}
