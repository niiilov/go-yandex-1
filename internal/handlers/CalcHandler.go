package handlers

import (
	"encoding/json"
	"github.com/niiilov/go-yandex-1/internal/models"
	"github.com/niiilov/go-yandex-1/pkg/calculation"
	"net/http"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed") //405
		return
	}

	var request models.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, "Bad Request") //400
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		errorResponse(w, http.StatusUnprocessableEntity, "Unprocessable Entity") //422
		return
	}

	res := models.Response{Result: result}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	json.NewEncoder(w).Encode(res)
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	err := models.Errors{Error: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}
