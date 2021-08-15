package controller

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func OKResponse(w http.ResponseWriter, payload interface{}) {
	JSON(w, http.StatusOK, payload)
}

func CreatedResponse(w http.ResponseWriter, payload interface{}) {
	JSON(w, http.StatusCreated, payload)
}

func BadRequestResponse(w http.ResponseWriter, err string) {
	JSON(w, http.StatusBadRequest, map[string]string{"error": err})
}

func NotFoundResponse(w http.ResponseWriter, err string) {
	JSON(w, http.StatusNotFound, map[string]string{"error": err})
}

func InternalServerErrorResponse(w http.ResponseWriter, err string) {
	JSON(w, http.StatusInternalServerError, map[string]string{"error": err})
}
