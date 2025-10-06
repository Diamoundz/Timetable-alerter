package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
    agendasCtrl "middleware/example/internal/controllers/agendas"
    alertsCtrl "middleware/example/internal/controllers/alerts"
    eventsCtrl "middleware/example/internal/controllers/events"
	"middleware/example/internal/controllers/users"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

    r.Route("/users", func(r chi.Router) { // route /users
		r.Get("/", users.GetUsers)            // GET /users
		r.Route("/{id}", func(r chi.Router) { // route /users/{id}
			r.Use(users.Context)      // Use Context method to get user ID
			r.Get("/", users.GetUser) // GET /users/{id}
		})
	})

    r.Route("/config_api", func(r chi.Router) {
        r.Route("/agendas", func(r chi.Router) {
            r.Get("/", agendasCtrl.GetAgendas)
            r.Post("/", agendasCtrl.PostAgendas)
            r.Route("/{id}", func(r chi.Router) {
                r.Use(agendasCtrl.Context)
                r.Get("/", agendasCtrl.GetAgenda)
                r.Put("/", agendasCtrl.PutAgenda)
                r.Delete("/", agendasCtrl.DeleteAgenda)
            })
        })
        r.Route("/alerts", func(r chi.Router) {
            r.Get("/", alertsCtrl.GetAlerts)
            r.Post("/", alertsCtrl.PostAlerts)
            r.Route("/{id}", func(r chi.Router) {
                r.Use(alertsCtrl.Context)
                r.Put("/", alertsCtrl.PutAlert)
                r.Delete("/", alertsCtrl.DeleteAlert)
            })
        })
    })

    r.Route("/timetable_api", func(r chi.Router) {
        r.Get("/events", eventsCtrl.GetEvents)
    })

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
    db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)

    // Config DB: agendas and alerts
    configDB, err := helpers.OpenConfigDB()
    if err != nil {
        logrus.Fatalf("error while opening config database : %s", err.Error())
    }
    configSchemes := []string{
        `CREATE TABLE IF NOT EXISTS agendas (
            id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
            uca_id INTEGER NOT NULL,
            name VARCHAR(255) NOT NULL
        );`,
        `CREATE TABLE IF NOT EXISTS alerts (
            id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
            email VARCHAR(255) NOT NULL,
            agenda_id VARCHAR(255) NOT NULL
        );`,
    }
    for _, scheme := range configSchemes {
        if _, err := configDB.Exec(scheme); err != nil {
            logrus.Fatalln("Could not generate config tables ! Error was : " + err.Error())
        }
    }
    helpers.CloseDB(configDB)
}
