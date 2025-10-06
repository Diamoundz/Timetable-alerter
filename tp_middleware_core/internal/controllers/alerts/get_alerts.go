package alerts

import (
    "encoding/json"
    "middleware/example/internal/helpers"
    service "middleware/example/internal/services/alerts"
    "net/http"
)

func GetAlerts(w http.ResponseWriter, _ *http.Request) {
    items, err := service.GetAll()
    if err != nil {
        body, status := helpers.RespondError(err)
        w.WriteHeader(status)
        if body != nil { _, _ = w.Write(body) }
        return
    }
    w.WriteHeader(http.StatusOK)
    out, _ := json.Marshal(items)
    _, _ = w.Write(out)
}


