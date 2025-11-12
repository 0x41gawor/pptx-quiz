package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{Port: port}
}
func (s *Server) Run() {
	router := mux.NewRouter()

	apiQuiz := NewHandlerQuiz()
	apiCompletions := NewHandlerCompletions()
	router.HandleFunc("/api/v1/quiz", makeHTTPHandleFunc(apiQuiz.handleBaseGET)).Methods("GET")
	router.HandleFunc("/api/v1/completions", makeHTTPHandleFunc(apiCompletions.handleBaseGET)).Methods("GET")
	router.HandleFunc("/api/v1/completions", makeHTTPHandleFunc(apiCompletions.handleBasePOST)).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"*"},
	})
	routerWithCors := c.Handler(router)

	log.Println("JSON API server running on port: ", s.Port)
	http.ListenAndServe(s.Port, routerWithCors)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// just calling the func
		err := f(w, r)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("error: %s", err.Error()))
		}
	}
}
