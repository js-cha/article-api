package controller

import (
	"encoding/json"
	"net/http"
)

const (
	BadRequestError     = "INVALID ID"
	NotFoundError       = "RESOURCE NOT FOUND"
	InternalServerError = "INTERNAL SERVER ERROR"
)

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func OKResponse(w http.ResponseWriter, payload interface{}) {
	JSON(w, http.StatusOK, payload)
}

func BadRequestResponse(w http.ResponseWriter) {
	JSON(w, http.StatusBadRequest, map[string]string{"error": BadRequestError})
}

func NotFoundResponse(w http.ResponseWriter) {
	JSON(w, http.StatusNotFound, map[string]string{"error": NotFoundError})
}

func InternalServerErrorResponse(w http.ResponseWriter) {
	JSON(w, http.StatusInternalServerError, map[string]string{"error": InternalServerError})
}
