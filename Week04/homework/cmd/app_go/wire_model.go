package main

import "fmt"

type Name string

func NewName(name string) Name {
	return Name(name)
}

type User struct {
	Name Name
}

func (u User) name() {
	fmt.Println(u.Name)
}

func NewUser(name Name) User {
	return User{
		Name: name,
	}
}

type MyWire struct {
	name Name
	User User
}

func NewMyWire(n Name, u User) MyWire {
	return MyWire{
		User: u,
		name: n,
	}
}
