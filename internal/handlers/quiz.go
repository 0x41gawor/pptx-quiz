package handlers

import (
	"log/slog"
	"net/http"

	"github.com/0x41gawor/pptx-quiz/internal/service"
)

type HandlerQuiz struct {
	s service.ServiceQuiz
}

func NewHandlerQuiz() *HandlerQuiz {
	return &HandlerQuiz{
		s: *service.NewServiceQuiz(),
	}
}

// handles "/quiz" path
func (h *HandlerQuiz) handleBaseGET(w http.ResponseWriter, r *http.Request) error {
	slog.Debug("")
	return WriteJSON(w, http.StatusOK, "quiz handler works")
}
