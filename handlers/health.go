package handlers

import (
    "encoding/json"
    "net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    encoder := json.NewEncoder(w)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    resp := map[string]interface{}{"status": "up"}

    if err := encoder.Encode(resp); err != nil {
        panic("unable to encode response")
    }
}
