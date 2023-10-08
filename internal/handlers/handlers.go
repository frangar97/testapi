package handlers

import "github.com/frangar97/testapi/internal/service"

type Handlers struct {
	secret   string
	services service.Service
}

type requestResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewHandlers(services service.Service, secret string) Handlers {
	return Handlers{secret: secret, services: services}
}
