package alerts

import (
    "encoding/json"
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
    service "middleware/example/internal/services/alerts"
    "net/http"
)

type putAlertBody struct {
    Email    string `json:"email"`
    AgendaId string `json:"agendaId"`
}

func PutAlert(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    alertId, _ := ctx.Value("alertId").(uuid.UUID)

    var body putAlertBody
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        resp, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "invalid request body"})
        w.WriteHeader(status)
        if resp != nil { _, _ = w.Write(resp) }
        return
    }
    agendaId, err := uuid.FromString(body.AgendaId)
    if err != nil {
        resp, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "invalid agendaId"})
        w.WriteHeader(status)
        if resp != nil { _, _ = w.Write(resp) }
        return
    }
    item, err := service.Update(alertId, body.Email, agendaId)
    if err != nil {
        resp, status := helpers.RespondError(err)
        w.WriteHeader(status)
        if resp != nil { _, _ = w.Write(resp) }
        return
    }
    w.WriteHeader(http.StatusOK)
    out, _ := json.Marshal(item)
    _, _ = w.Write(out)
}


