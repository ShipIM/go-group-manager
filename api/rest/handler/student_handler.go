package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

var studentsPathPrefix = "/api/v1/students/"

func (h *Handler) handleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createStudent(w, r)
	case "GET":
		h.getStudents(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	created, err := h.service.CreateStudent(student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%d", created.Id)
}

func (h *Handler) getStudents(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getStudent(w, r)
	case "PUT":
		h.updateStudent(w, r)
	case "DELETE":
		h.deleteStudent(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, studentsPathPrefix)
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	student, err := h.service.GetStudentById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(student)

	fmt.Fprintf(w, "%s", reqBodyBytes.String())
}

func (h *Handler) updateStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	idString := strings.TrimPrefix(r.URL.Path, studentsPathPrefix)
	student.Id, err = strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = h.service.UpdateStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *Handler) deleteStudent(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, studentsPathPrefix)
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = h.service.DeleteStudentById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
