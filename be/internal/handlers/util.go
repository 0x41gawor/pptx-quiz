package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Res string `json:"res"`
}

func NewResponse(res string) *Response {
	return &Response{
		Res: res,
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// GetQueryString pobiera parametr typu string z query lub zwraca domyślną wartość.
func GetQueryString(r *http.Request, key string, defaultVal string) string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// GetQueryInt pobiera parametr typu int z query lub zwraca domyślną wartość,
// jeśli parametr nie istnieje lub nie można go sparsować.
func GetQueryInt(r *http.Request, key string, defaultVal int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}
