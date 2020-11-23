package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// HTTPSuccess represents a successful API operation, can be marshalled into JSON
type HTTPSuccess struct {
	Success bool   `json:"success"`
	String  string `json:"string"`
	Answer  int    `json:"answer"`
}

// Success constructs, marshals then writes an HTTP success into the response in JSON
func Success(w http.ResponseWriter, success bool, str string, answer int) {
	json, err := json.Marshal(&HTTPSuccess{
		Success: success,
		String:  str,
		Answer:  answer,
	})
	if err != nil {
		// Should not occur, panic
		log.Panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
