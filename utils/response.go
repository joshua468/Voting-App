package utils

import (
    "encoding/json"
    "net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := map[string]interface{}{
        "message": message,
        "data":    data,
    }
    json.NewEncoder(w).Encode(response)
}
