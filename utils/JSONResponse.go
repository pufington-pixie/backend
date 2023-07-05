package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pufington-pixie/haver/pkg/models"
)

func HandleError(w http.ResponseWriter, err error, status int, message string) {
	log.Print(err)

	response := models.Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}

	SendJSONResponse(w, response, status)
}

func SendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Failed to encode JSON response:", err)
	}
}
