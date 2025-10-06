package events

import (
    "encoding/json"
    "middleware/example/internal/helpers"
    "net/http"
)

// For now, return an empty list; repository/service can be added later
func GetEvents(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    out, _ := json.Marshal([]interface{}{})
    _, _ = w.Write(out)
}


