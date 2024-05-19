package http

import (
	"net/http"

	rate_dto "exchange_rate/pkg/usecase/rate/dto"

	"github.com/gin-gonic/gin"
)

func (h *HTTPController) CurrencyController(router *gin.Engine) {
	currencyGroup := router.Group("/rate")

	currencyGroup.Use()
	{
		currencyGroup.GET("/",
			func(c *gin.Context) {
				ctx := c.Request.Context()

				response, err := h.services.GetRateByValCode(ctx, &rate_dto.GetRateByValCodeDto{
					RateValCode: rate_dto.RateValCode{ValCode: h.baseValCode},
				})
				if err != nil {
					c.JSON(400, nil)
					return
				}

				c.JSON(http.StatusOK, response.Rate)
			},
		)
	}
}
