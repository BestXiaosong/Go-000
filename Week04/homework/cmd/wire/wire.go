//+build wireinject
package main

import (
	"github.com/google/wire"
)

func InitController(conf int, name string) *Controller {
	wire.Build(NewLogin, NewController, NewModel)
	return &Controller{}
}
