package http

import (
	"exchange_rate/pkg/domain/vo"
	"exchange_rate/pkg/packages/errors"
	"exchange_rate/pkg/packages/validator"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HTTPController struct {
	services    Services
	handlers    *gin.Engine
	v           *validator.Validator
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
		services:    services,
		handlers:    gin.Default(),
		v:           validator.NewValidator(),
		baseValCode: baseValCode,
	}, nil
}

func (h *HTTPController) InitControllers() *gin.Engine {
	h.CurrencyController(h.handlers)
	h.EmailController(h.handlers)

	if err := h.v.Register("customUUID", vo.ValidateUUID); err != nil {
		logrus.Fatalf("error add customUUID validation func: %v", err)
	}

	return h.handlers
}
