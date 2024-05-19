package http

import (
	"exchange_rate/pkg/domain/vo"
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
				// email := domain.NewUserEmail(c.PostForm("email"))

				// if err := email.Validate(); err != nil {
				// 	c.JSON(http.StatusBadRequest, err.Error())
				// 	return
				// }

				if err := h.services.CreateUser(ctx, &user_dto.CreateUserDto{
					UserEmail: user_dto.UserEmail{
						Mail: vo.Email(c.PostForm("email")),
					},
				}); err != nil {
					// if err == domain.ErrAlreadyExist {
					// c.JSON(http.StatusConflict, nil)
					c.JSON(err.GetCode(), nil)
					return
					// }
					// c.JSON(http.StatusInternalServerError, nil)
					// return
				}

				c.JSON(http.StatusOK, nil)
			},
		)
	}
}
