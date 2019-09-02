package main

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     db.UserDb
	Ctrlr  controller.Controller
}

func (a *App) Initialize(dbname string) {}

func (a *App) Run(addr string) {}
