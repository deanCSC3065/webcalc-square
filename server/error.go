package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HTTPError represents when an error has occurred with relevant info, can be marshalled into JSON
type HTTPError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// HandleNotFound handles 404 errors, providing an HTTP error in JSON
func HandleNotFound() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Error(w, false, fmt.Sprintf("Endpoint '%s' not found.", r.URL.Path), http.StatusNotFound)
	})
}

// HandleMethodNotAllowed handles 405 errors, providing an HTTP error in JSON
func HandleMethodNotAllowed() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Error(w, false, fmt.Sprintf("Method '%s' not allowed for endpoint '%s'.", r.Method, r.URL.Path), http.StatusMethodNotAllowed)
	})
}

// Error constructs, marshals then writes an HTTP error into the response in JSON
func Error(w http.ResponseWriter, success bool, message string, code int) {
	json, err := json.Marshal(&HTTPError{
		Success: success,
		Message: message,
		Code:    code,
	})
	if err != nil {
		// Should not occur, panic
		log.Panic(err)
		return
	}
	w.WriteHeader(code)
	w.Write(json)
}
