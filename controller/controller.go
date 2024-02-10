package controller

type Controller struct {
	Hello
	User
	Card
	Poem
}

func New() *Controller {
	Controller := &Controller{}
	return Controller
}
