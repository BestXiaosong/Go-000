package inic

type Controller struct {
	Model Model
	Logic Logic
}

func NewController(model Model, logic Logic) *Controller {
	return &Controller{
		Model: model,
		Logic: logic,
	}
}
