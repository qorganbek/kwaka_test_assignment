package handler

import (
	"github.com/qorganbek/kwaka_test_assignment/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}
