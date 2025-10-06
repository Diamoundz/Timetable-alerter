package agendas

import (
    "encoding/json"
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
    service "middleware/example/internal/services/agendas"
    "net/http"
)

type putAgendaBody struct {
    UcaId interface{} `json:"ucaId"`
    Name  string      `json:"name"`
}

func PutAgenda(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    agendaId, _ := ctx.Value("agendaId").(uuid.UUID)

    var body putAgendaBody
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        resp, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "invalid request body"})
        w.WriteHeader(status)
        if resp != nil { _, _ = w.Write(resp) }
        return
    }

    var ucaId int
    switch v := body.UcaId.(type) {
    case float64:
        ucaId = int(v)
    case string:
        // ignore parse errors to keep behavior aligned with front
    }

    item, err := service.Update(agendaId, ucaId, body.Name)
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


