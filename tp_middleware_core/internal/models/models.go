package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

// Agenda represents a remote UCA agenda subscription stored in Config API
type Agenda struct {
    Id    *uuid.UUID `json:"id"`
    UcaId int        `json:"ucaId"`
    Name  string     `json:"name"`
}

// Alert represents a notification recipient bound to an agenda (Config API)
type Alert struct {
    Id       *uuid.UUID `json:"id"`
    Email    string     `json:"email"`
    AgendaId *uuid.UUID `json:"agendaId"`
}

// Event is a read-only entity exposed by the Timetable API
type Event struct {
    Id         *uuid.UUID  `json:"id"`
    AgendaIds  []uuid.UUID `json:"agendaIds"`
    Uid        string      `json:"uid"`
    Description string     `json:"description"`
    Name       string      `json:"name"`
    Start      string      `json:"start"`
    End        string      `json:"end"`
    Location   string      `json:"location"`
    LastUpdate string      `json:"lastUpdate"`
}
