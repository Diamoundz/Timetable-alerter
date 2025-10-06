package agendas

import (
    "encoding/json"
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    service "middleware/example/internal/services/agendas"
    "net/http"
)

// GetAgenda
// @Tags         agendas
// @Summary      Get an agenda.
// @Description  Get an agenda.
// @Param        id            path      string  true  "Agenda UUID formatted ID"
// @Success      200            {object}  models.Agenda
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /config_api/agendas/{id} [get]
func GetAgenda(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    agendaId, _ := ctx.Value("agendaId").(uuid.UUID)
    item, err := service.GetById(agendaId)
    if err != nil {
        body, status := helpers.RespondError(err)
        w.WriteHeader(status)
        if body != nil {
            _, _ = w.Write(body)
        }
        return
    }
    w.WriteHeader(http.StatusOK)
    body, _ := json.Marshal(item)
    _, _ = w.Write(body)
}


