package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

var items = []string{
	"Apple", "Banana", "Cherry", "Date", "Elderberry",
	"Fig", "Grape", "Honeydew", "Kiwi", "Lemon",
	"Mango", "Nectarine", "Orange", "Papaya", "Quince",
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	query := strings.ToLower(req.Value)
	filtered := []string{}
	for _, item := range items {
		if strings.Contains(strings.ToLower(item), query) {
			filtered = append(filtered, item)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/options", optionsHandler)
	mux.HandleFunc("/api/submit", submitHandler)

	log.Println("Backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}
