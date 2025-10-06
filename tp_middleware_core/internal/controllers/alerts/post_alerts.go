package alerts

import (
    "encoding/json"
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
    service "middleware/example/internal/services/alerts"
    "net/http"
)

type postAlertBody struct {
    Email    string `json:"email"`
    AgendaId string `json:"agendaId"`
}

func PostAlerts(w http.ResponseWriter, r *http.Request) {
    var body postAlertBody
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
    item, err := service.Create(body.Email, agendaId)
    if err != nil {
        resp, status := helpers.RespondError(err)
        w.WriteHeader(status)
        if resp != nil { _, _ = w.Write(resp) }
        return
    }
    w.WriteHeader(http.StatusCreated)
    out, _ := json.Marshal(item)
    _, _ = w.Write(out)
}


