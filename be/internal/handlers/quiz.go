package handlers

import (
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

// GET /quiz?howMany=10&lang=pl
func (h *HandlerQuiz) handleBaseGET(w http.ResponseWriter, r *http.Request) error {
	howMany := GetQueryInt(r, "howMany", 10)
	lang := GetQueryString(r, "lang", "pl")

	// langu musi być albo "pl" albo "en"
	if lang != "pl" && lang != "en" {
		lang = "pl"
	}

	quiz := h.s.Get(lang, howMany)
	return WriteJSON(w, http.StatusOK, quiz)
}
