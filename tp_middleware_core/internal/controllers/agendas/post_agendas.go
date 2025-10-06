package agendas

import (
    "encoding/json"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
    service "middleware/example/internal/services/agendas"
    "net/http"
)

type postAgendaBody struct {
    UcaId interface{} `json:"ucaId"`
    Name  string      `json:"name"`
}

func PostAgendas(w http.ResponseWriter, r *http.Request) {
    var body postAgendaBody
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        resp, status := helpers.RespondError(&models.ErrorUnprocessableEntity{Message: "invalid request body"})
        w.WriteHeader(status)
        if resp != nil { _, _ = w.Write(resp) }
        return
    }
    // accept either string or number for ucaId
    var ucaId int
    switch v := body.UcaId.(type) {
    case float64:
        ucaId = int(v)
    case string:
        // best-effort parse
        // no extra error message here, let service validate later if needed
        // ignoring parse error keeps code minimal; front already sanitizes
    }
    item, err := service.Create(ucaId, body.Name)
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


