package alerts

import (
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    service "middleware/example/internal/services/alerts"
    "net/http"
)

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    alertId, _ := ctx.Value("alertId").(uuid.UUID)
    if err := service.Delete(alertId); err != nil {
        body, status := helpers.RespondError(err)
        w.WriteHeader(status)
        if body != nil { _, _ = w.Write(body) }
        return
    }
    w.WriteHeader(http.StatusNoContent)
}


