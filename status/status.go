package status

import (
	"encoding/json"
	"net/http"
)

type emptyModel struct {
	Status *Status `json:"status"`
}

type ModelError struct {

	Code string `json:"code"`

	Message string `json:"message"`
}

type Status struct {

	Ok bool `json:"ok"`

	Err *ModelError `json:"err,omitempty"`
}

func WriteEmptyError(w http.ResponseWriter, code string, message string) {
	errorModel := ModelError{code, message}
	statusModel := Status{false, &errorModel}
	writeStatus(w, statusModel)
}

func WriteEmptyOk(w http.ResponseWriter) {
	statusModel := Status{true, nil}
	writeStatus(w, statusModel)
}

func writeStatus(w http.ResponseWriter, statusModel Status) {
	model := emptyModel{&statusModel}
	status, err :=json.Marshal(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(status)
}
