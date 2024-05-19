package controllers

import (
	_http "exchange_rate/pkg/controllers/http"
	"exchange_rate/pkg/packages/errors"
	"log"
	"net/http"
)

type Controllers struct {
	HTTP      *_http.HTTPController
	serverURL string
}

func NewControllers(services _http.Services, serverURL, basicValCode string) (*Controllers, *errors.Error) {
	http, err := _http.NewHttpControllers(services, basicValCode)
	if err != nil {
		return nil, err
	}

	return &Controllers{
		HTTP:      http,
		serverURL: serverURL,
	}, nil
}

func (c *Controllers) Start() {
	handlers := c.HTTP.InitControllers()

	go listenAndServe(c.serverURL, handlers)
}

func listenAndServe(url string, handlers http.Handler) {
	err := http.ListenAndServe(url, handlers)
	if err != nil {
		log.Panic(err)
	}
}
