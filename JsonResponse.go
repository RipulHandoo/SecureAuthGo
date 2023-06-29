package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
    data, err := json.Marshal(payload)
    if err != nil {
        log.Fatal(err)
        return
    }

    w.Header().Add("Content-Type", "application/json") // Corrected header field
    w.WriteHeader(code)
    w.Write(data)
}
func ResponseWithError(w http.ResponseWriter, code int, err error) {
	if err != nil {
		ResponseWithJson(w, code, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	} else {
		// If err is nil, handle it as a generic internal server error
		ResponseWithJson(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: "Internal Server Error",
		})
	}
}
