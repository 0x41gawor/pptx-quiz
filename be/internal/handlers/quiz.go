package handlers

import (
	"encoding/csv"
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

func (h *HandlerQuiz) handleFileGET(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="quiz.csv"`)

	writer := csv.NewWriter(w)
	writer.Comma = ';'

	for _, row := range h.s.Records[0:] {
		writer.Write(row)
	}
	writer.Flush()

	return nil
}

// POST /quiz/file
func (h *HandlerQuiz) handleFilePOST(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20) // max 10MB
	if err != nil {
		print("error parsing multipart form: ", err)
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		print("error retrieving the file: ", err)
		return WriteJSON(w, http.StatusBadRequest, err)
	}
	defer file.Close()

	// wczytujemy CSV do service
	if err := h.s.LoadCSV(file); err != nil {
		print("erro loading csv")
		return WriteJSON(w, http.StatusBadRequest, err)
	}

	return WriteJSON(w, http.StatusOK, map[string]string{
		"status": "csv loaded",
	})
}
