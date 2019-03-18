package Controller

import (
	"encoding/json"
	"net/http"

	"../Model"
)

func Test(w http.ResponseWriter, r *http.Request) {
	var (
		status Model.Status
	)

	w.Header().Set("Content-Type", "application/json")

	status.Status = 200
	status.Comment = "tes"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
