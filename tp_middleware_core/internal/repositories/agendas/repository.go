package agendas

import (
    "database/sql"
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetAll() ([]models.Agenda, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    rows, err := db.Query("SELECT id, uca_id, name FROM agendas")
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }

    items := []models.Agenda{}
    for rows.Next() {
        var idStr string
        var item models.Agenda
        if err = rows.Scan(&idStr, &item.UcaId, &item.Name); err != nil {
            _ = rows.Close()
            return nil, err
        }
        if parsed, e := uuid.FromString(idStr); e == nil {
            item.Id = &parsed
        }
        items = append(items, item)
    }
    _ = rows.Close()
    return items, nil
}

func GetById(id uuid.UUID) (*models.Agenda, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    row := db.QueryRow("SELECT id, uca_id, name FROM agendas WHERE id=?", id.String())
    helpers.CloseDB(db)

    var idStr string
    var item models.Agenda
    if err = row.Scan(&idStr, &item.UcaId, &item.Name); err != nil {
        return nil, err
    }
    if parsed, e := uuid.FromString(idStr); e == nil {
        item.Id = &parsed
    }
    return &item, nil
}

func Create(ucaId int, name string) (*models.Agenda, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    newId, _ := uuid.NewV4()
    _, err = db.Exec("INSERT INTO agendas (id, uca_id, name) VALUES (?, ?, ?)", newId.String(), ucaId, name)
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }
    return GetById(newId)
}

func Update(id uuid.UUID, ucaId int, name string) (*models.Agenda, error) {
    db, err := helpers.OpenConfigDB()
    if err != nil {
        return nil, err
    }
    res, err := db.Exec("UPDATE agendas SET uca_id=?, name=? WHERE id=?", ucaId, name, id.String())
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
    res, err := db.Exec("DELETE FROM agendas WHERE id=?", id.String())
    helpers.CloseDB(db)
    if err != nil {
        return err
    }
    if n, _ := res.RowsAffected(); n == 0 {
        return sql.ErrNoRows
    }
    return nil
}


