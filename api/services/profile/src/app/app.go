package app

type App struct {}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
	return a
}

func (a *App) Serve() {}