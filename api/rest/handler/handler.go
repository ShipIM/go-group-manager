package handler

import (
	"net/http"

	"github.com/ShipIM/go-group-manager/internal/service"
)

type Handler struct {
	service.GroupService
	service.StudentService
}

func NewHandler(groupService service.GroupService, studentService service.StudentService) *Handler {
	return &Handler{
		groupService,
		studentService,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/groups", http.HandlerFunc(h.handleGroups))
	mux.Handle("/api/v1/groups/", http.HandlerFunc(h.handleGroup))

	mux.Handle("/api/v1/students", http.HandlerFunc(h.handleStudents))
	mux.Handle("/api/v1/students/", http.HandlerFunc(h.handleStudent))

	return mux
}
