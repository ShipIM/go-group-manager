package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ShipIM/go-group-manager/sender/internal/domain/entity"
)

var groupPathPrefix = "/api/v1/groups/"

func (h *Handler) handleGroups(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createGroup(w, r)
	case http.MethodGet:
		h.getGroups(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createGroup(w http.ResponseWriter, r *http.Request) {
	var group entity.Group

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.CreateGroup(group)
	if err != nil {
		http.Error(w, "failed to create group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		http.Error(w, "failed to encode group to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getGroups(w http.ResponseWriter, _ *http.Request) {
	groups, err := h.GetGroups()
	if err != nil {
		http.Error(w, "failed to get groups: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(groups)
	if err != nil {
		http.Error(w, "failed to encode groups to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleGroup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getGroup(w, r)
	case http.MethodPut:
		h.updateGroup(w, r)
	case http.MethodDelete:
		h.deleteGroup(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getGroup(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, groupPathPrefix)

	group, err := h.GetGroupByName(name)
	if err != nil {
		http.Error(w, "failed to get group: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		http.Error(w, "failed to encode group to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateGroup(w http.ResponseWriter, r *http.Request) {
	var group entity.Group

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	group.Name = strings.TrimPrefix(r.URL.Path, groupPathPrefix)

	err = h.UpdateGroup(group)
	if err != nil {
		http.Error(w, "failed to update group: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		http.Error(w, "failed to encode group to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, groupPathPrefix)

	err := h.DeleteGroupByName(name)
	if err != nil {
		http.Error(w, "failed to delete group: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
