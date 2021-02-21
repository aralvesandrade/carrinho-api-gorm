package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	contentTypeConst = "Content-Type"
	applicationJSON  = "application/json"
)

//NewHealthCheckHandler ...
func NewHealthCheckHandler(e *mux.Router) {
	s := e.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/health", HealthCheck).Methods(http.MethodGet)
}

//HealthCheck ...
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(contentTypeConst, applicationJSON)

	response, _ := json.Marshal(map[string]interface{}{"status": "up"})

	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
