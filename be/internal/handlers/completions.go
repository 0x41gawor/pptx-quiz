package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/0x41gawor/pptx-quiz/internal/service"
	"github.com/0x41gawor/pptx-quiz/internal/service/model"
)

// HandlerCompletions obsługuje endpointy związane z ukończeniami kursów.
type HandlerCompletions struct {
	s service.ServiceCompletions
}

func NewHandlerCompletions() *HandlerCompletions {
	return &HandlerCompletions{
		s: *service.NewServiceCompletions(),
	}
}

// GET /completions
func (h *HandlerCompletions) handleBaseGET(w http.ResponseWriter, r *http.Request) error {
	completions, err := h.s.List()
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, "database query failed")
	}

	return WriteJSON(w, http.StatusOK, completions)
}

// POST /completions
func (h *HandlerCompletions) handleBasePOST(w http.ResponseWriter, r *http.Request) error {
	var c model.Completion
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		return WriteJSON(w, http.StatusBadRequest, "invalid JSON body")
	}

	newC, err := h.s.Create(&c)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, "database insert failed")
	}

	return WriteJSON(w, http.StatusCreated, newC)
}
