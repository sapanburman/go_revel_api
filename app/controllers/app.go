package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}


//type User struct {
//	*revel.Controller
//}
//
//var users = []models.User{
//	models.User{1, "Sapan", "sapanbarman09@gmail.com", "Node.js Developer"},
//	models.User{2, "Rohan", "rohan@gmail.com", "Python Developer"},
//	models.User{3, "Mohan", "mohan@gmail.com", "Java Developer"},
//}

func (c App) Index() revel.Result {
	return c.Render()
}

//func(c User) Index() revel.Result{
//	return c.RenderJSON(users)
//}