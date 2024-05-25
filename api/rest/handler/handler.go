package handler

import (
	"net/http"

	"github.com/ShipIM/go-group-manager/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/groups", http.HandlerFunc(h.handleGroups))
	mux.Handle("/api/v1/groups/", http.HandlerFunc(h.handleGroup))

	mux.Handle("/api/v1/students", http.HandlerFunc(h.handleStudents))
	mux.Handle("/api/v1/students/", http.HandlerFunc(h.handleStudent))

	return mux
}
