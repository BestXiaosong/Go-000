package main

import (
	"Go-001/Week04/homework/cmd/inic"
	"github.com/google/wire"
)

func InitController() *inic.Controller {

	wire.Build(inic.NewLogin, inic.NewModel, inic.NewController)
	return &inic.Controller{}
}
