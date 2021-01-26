package main

type Logic struct {
	Conf int
}

func NewLogin(name int) *Logic {
	return &Logic{Conf: name}
}

type Model struct {
	Name string
}

func NewModel(name string) *Model {
	return &Model{Name: name}
}

type Controller struct {
	Model *Model
	Logic *Logic
}

func NewController(model *Model, logic *Logic) *Controller {
	return &Controller{
		Model: model,
		Logic: logic,
	}
}
