package web

import "backend-kata/config"

type App struct {
	Server *Server
}

func NewApplication() *App {
	cfg := config.ReadConfig()
	container := NewContainer(cfg)
	server := NewServer(cfg)
	server.RegisterRoutes(container)
	return &App{
		Server: server,
	}
}

func (a *App) Run() {
	a.Server.Run()
}
