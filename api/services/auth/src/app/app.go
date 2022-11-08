package app

import (
	"log"

	"github.com/ssibrahimbas/samchat-v3.auth/src/config"
	cnf "github.com/ssibrahimbas/samchat-v3.shared/pkg/config"
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/http"
)

type App struct {
	Cnf  *config.App
	Http *http.Client
}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
	return a
}

func (a *App) Serve() {
	log.Fatal(a.Http.Listen(a.Cnf.Host + a.Cnf.Port))
}

func (a *App) loadConfig() {
	cnfg := config.App{}
	cnf.LoadConfig(&cnfg)
	a.Cnf = &cnfg
}
