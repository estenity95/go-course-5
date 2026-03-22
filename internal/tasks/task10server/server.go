package task10server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type SumResponse struct {
	Result int `json:"result"`
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/sum", sumHandler)
	return mux
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, "invalid parameter a", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		http.Error(w, "invalid parameter b", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(SumResponse{Result: a + b}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
