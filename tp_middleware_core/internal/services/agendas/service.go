package agendas

import (
    "database/sql"
    "fmt"
    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/models"
    repo "middleware/example/internal/repositories/agendas"
)

func GetAll() ([]models.Agenda, error) {
    items, err := repo.GetAll()
    if err != nil {
        logrus.Errorf("error retrieving agendas : %s", err.Error())
        return nil, &models.ErrorGeneric{Message: "Something went wrong while retrieving agendas"}
    }
    return items, nil
}

func GetById(id uuid.UUID) (*models.Agenda, error) {
    item, err := repo.GetById(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, &models.ErrorNotFound{Message: "agenda not found"}
        }
        logrus.Errorf("error retrieving agenda %s : %s", id.String(), err.Error())
        return nil, &models.ErrorGeneric{Message: fmt.Sprintf("Something went wrong while retrieving agenda %s", id.String())}
    }
    return item, nil
}

func Create(ucaId int, name string) (*models.Agenda, error) {
    item, err := repo.Create(ucaId, name)
    if err != nil {
        logrus.Errorf("error creating agenda : %s", err.Error())
        return nil, &models.ErrorGeneric{Message: "Something went wrong while creating agenda"}
    }
    return item, nil
}

func Update(id uuid.UUID, ucaId int, name string) (*models.Agenda, error) {
    item, err := repo.Update(id, ucaId, name)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, &models.ErrorNotFound{Message: "agenda not found"}
        }
        logrus.Errorf("error updating agenda %s : %s", id.String(), err.Error())
        return nil, &models.ErrorGeneric{Message: fmt.Sprintf("Something went wrong while updating agenda %s", id.String())}
    }
    return item, nil
}

func Delete(id uuid.UUID) error {
    if err := repo.Delete(id); err != nil {
        if err == sql.ErrNoRows {
            return &models.ErrorNotFound{Message: "agenda not found"}
        }
        logrus.Errorf("error deleting agenda %s : %s", id.String(), err.Error())
        return &models.ErrorGeneric{Message: fmt.Sprintf("Something went wrong while deleting agenda %s", id.String())}
    }
    return nil
}


