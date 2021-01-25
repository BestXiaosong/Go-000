package inic

type Model struct {
	name string
}

func NewModel() *Model {
	return &Model{name: "xiaosong"}
}
