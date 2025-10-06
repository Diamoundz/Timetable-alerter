package alerts

import (
    "database/sql"
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetAll() ([]models.Alert, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    rows, err := db.Query("SELECT id, email, agenda_id FROM alerts")
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }
    items := []models.Alert{}
    for rows.Next() {
        var idStr, agendaIdStr string
        var item models.Alert
        if err = rows.Scan(&idStr, &item.Email, &agendaIdStr); err != nil {
            _ = rows.Close()
            return nil, err
        }
        if parsed, e := uuid.FromString(idStr); e == nil {
            item.Id = &parsed
        }
        if parsed, e := uuid.FromString(agendaIdStr); e == nil {
            item.AgendaId = &parsed
        }
        items = append(items, item)
    }
    _ = rows.Close()
    return items, nil
}

func GetById(id uuid.UUID) (*models.Alert, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    row := db.QueryRow("SELECT id, email, agenda_id FROM alerts WHERE id=?", id.String())
    helpers.CloseDB(db)

    var idStr, agendaIdStr string
    var item models.Alert
    if err = row.Scan(&idStr, &item.Email, &agendaIdStr); err != nil {
        return nil, err
    }
    if parsed, e := uuid.FromString(idStr); e == nil {
        item.Id = &parsed
    }
    if parsed, e := uuid.FromString(agendaIdStr); e == nil {
        item.AgendaId = &parsed
    }
    return &item, nil
}

func Create(email string, agendaId uuid.UUID) (*models.Alert, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    newId, _ := uuid.NewV4()
    _, err = db.Exec("INSERT INTO alerts (id, email, agenda_id) VALUES (?, ?, ?)", newId.String(), email, agendaId.String())
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }
    return GetById(newId)
}

func Update(id uuid.UUID, email string, agendaId uuid.UUID) (*models.Alert, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    res, err := db.Exec("UPDATE alerts SET email=?, agenda_id=? WHERE id=?", email, agendaId.String(), id.String())
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }
    if n, _ := res.RowsAffected(); n == 0 {
        return nil, sql.ErrNoRows
    }
    return GetById(id)
}

func Delete(id uuid.UUID) error {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return err
    }
    res, err := db.Exec("DELETE FROM alerts WHERE id=?", id.String())
    helpers.CloseDB(db)
    if err != nil {
        return err
    }
    if n, _ := res.RowsAffected(); n == 0 {
        return sql.ErrNoRows
    }
    return nil
}


