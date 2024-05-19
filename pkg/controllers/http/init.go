package http

import (
	"exchange_rate/pkg/packages/errors"

	"github.com/gin-gonic/gin"
)

type HTTPController struct {
	services    Services
	handlers    *gin.Engine
	baseValCode string
}

func NewHttpControllers(
	services Services,
	baseValCode string,
) (*HTTPController, *errors.Error) {
	if services == nil {
		return nil, ErrInitHTTPController
	}

	return &HTTPController{
		services: services,
		handlers: gin.Default(),
	}, nil
}

func (h *HTTPController) InitControllers() *gin.Engine {
	h.CurrencyController(h.handlers)
	h.EmailController(h.handlers)

	return h.handlers
}
