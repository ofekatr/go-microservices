package main

import (
	"encoding/json"
	"net/http"
)

type jsonResopnse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (a *App) Broker(w http.ResponseWriter, r *http.Request) {
	payload := &jsonResopnse{
		Error:   false,
		Message: "Hit the borker",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
