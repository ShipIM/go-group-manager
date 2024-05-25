package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

var studentsPathPrefix = "/api/v1/students/"

func (h *Handler) handleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createStudent(w, r)
	case http.MethodGet:
		h.getStudents(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.CreateStudent(student)
	if err != nil {
		http.Error(w, "failed to create student: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		http.Error(w, "failed to encode student to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getStudents(w http.ResponseWriter, _ *http.Request) {
	students, err := h.GetStudents()
	if err != nil {
		http.Error(w, "failed to get students: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, "failed to encode students to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getStudent(w, r)
	case http.MethodPut:
		h.updateStudent(w, r)
	case http.MethodDelete:
		h.deleteStudent(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, studentsPathPrefix)
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid student ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	student, err := h.GetStudentById(id)
	if err != nil {
		http.Error(w, "failed to get student: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, "failed to encode student to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	idString := strings.TrimPrefix(r.URL.Path, studentsPathPrefix)
	student.Id, err = strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid student ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UpdateStudent(student)
	if err != nil {
		http.Error(w, "failed to update student: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, "failed to encode student to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deleteStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, studentsPathPrefix)
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid student ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.DeleteStudentById(id)
	if err != nil {
		http.Error(w, "failed to delete student: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
