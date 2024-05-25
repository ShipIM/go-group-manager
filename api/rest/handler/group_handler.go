package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

var groupPathPrefix = "/api/v1/groups/"

func (h *Handler) handleGroups(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createGroup(w, r)
	case "GET":
		h.getGroups(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createGroup(w http.ResponseWriter, r *http.Request) {
	var group entity.Group

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	created, err := h.service.CreateGroup(group)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", created.Name)
}

func (h *Handler) getGroups(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleGroup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getGroup(w, r)
	case "PUT":
		h.updateGroup(w, r)
	case "DELETE":
		h.deleteGroup(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getGroup(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, groupPathPrefix)

	group, err := h.service.GetGroupByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(group)

	fmt.Fprintf(w, "%s", reqBodyBytes.String())
}

func (h *Handler) updateGroup(w http.ResponseWriter, r *http.Request) {
	var group entity.Group

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	group.Name = strings.TrimPrefix(r.URL.Path, groupPathPrefix)

	err = h.service.UpdateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *Handler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, groupPathPrefix)

	err := h.service.DeleteGroupByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
