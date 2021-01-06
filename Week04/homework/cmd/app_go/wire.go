package main

import (
	"github.com/google/wire"
)

func InitializeSpeaker(name string) User {
	wire.Build(NewUser, NewName)
	return User{}
}
