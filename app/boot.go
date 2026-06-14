package app

import (
	"net/http"

	routes "github.com/SQU1DMAN6/ningchat"
	"github.com/SQU1DMAN6/ningchat/config"
	"github.com/go-chi/chi/v5"
)

func BootApp(isProduction bool) {
	r := chi.NewRouter()
	config.ConnectDatabase(isProduction)
	// config.InitSession()

	// r.Use(config.GetSessionManager().LoadAndSave)

	// RegisterMiddlewares(r)
	routes.RegisterRoutes(r)

	http.ListenAndServe(":6060", r)
}
