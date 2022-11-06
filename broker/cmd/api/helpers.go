package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type jsonResopnse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (a *App) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(data); err != nil {
		return errors.Wrap(err, "cannot decode json")
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return errors.Wrap(err, "request body must only contain a single JSON object")
	}

	return nil
}

func (a *App) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "cannot encode json")
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return errors.Wrap(err, "cannot write output json")
	}

	return nil
}

func (a *App) errorJSON(w http.ResponseWriter, e error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := jsonResopnse{
		Error:   true,
		Message: e.Error(),
	}

	if err := a.writeJSON(w, statusCode, payload); err != nil {
		return errors.Wrap(err, "cannot write error json")
	}

	return nil
}
