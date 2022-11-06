package main

import (
	"net/http"
)

func (a *App) Broker(w http.ResponseWriter, r *http.Request) {
	payload := &jsonResopnse{
		Error:   false,
		Message: "Hit the borker",
	}

	a.writeJSON(w, http.StatusAccepted, payload, http.Header{
		"Content-Type": []string{"application/json"},
	})
}
