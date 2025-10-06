package alerts

import (
    "database/sql"
    "fmt"
    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/models"
    repo "middleware/example/internal/repositories/alerts"
)

func GetAll() ([]models.Alert, error) {
    items, err := repo.GetAll()
    if err != nil {
        logrus.Errorf("error retrieving alerts : %s", err.Error())
        return nil, &models.ErrorGeneric{Message: "Something went wrong while retrieving alerts"}
    }
    return items, nil
}

func GetById(id uuid.UUID) (*models.Alert, error) {
    item, err := repo.GetById(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, &models.ErrorNotFound{Message: "alert not found"}
        }
        logrus.Errorf("error retrieving alert %s : %s", id.String(), err.Error())
        return nil, &models.ErrorGeneric{Message: fmt.Sprintf("Something went wrong while retrieving alert %s", id.String())}
    }
    return item, nil
}

func Create(email string, agendaId uuid.UUID) (*models.Alert, error) {
    item, err := repo.Create(email, agendaId)
    if err != nil {
        logrus.Errorf("error creating alert : %s", err.Error())
        return nil, &models.ErrorGeneric{Message: "Something went wrong while creating alert"}
    }
    return item, nil
}

func Update(id uuid.UUID, email string, agendaId uuid.UUID) (*models.Alert, error) {
    item, err := repo.Update(id, email, agendaId)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, &models.ErrorNotFound{Message: "alert not found"}
        }
        logrus.Errorf("error updating alert %s : %s", id.String(), err.Error())
        return nil, &models.ErrorGeneric{Message: fmt.Sprintf("Something went wrong while updating alert %s", id.String())}
    }
    return item, nil
}

func Delete(id uuid.UUID) error {
    if err := repo.Delete(id); err != nil {
        if err == sql.ErrNoRows {
            return &models.ErrorNotFound{Message: "alert not found"}
        }
        logrus.Errorf("error deleting alert %s : %s", id.String(), err.Error())
        return &models.ErrorGeneric{Message: fmt.Sprintf("Something went wrong while deleting alert %s", id.String())}
    }
    return nil
}


