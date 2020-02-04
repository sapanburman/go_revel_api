package controllers

import (

	"github.com/revel/revel"
	"log"
	"revelapi/app/models"
	"time"
)


type Employee struct {
	*revel.Controller
}

func (c Employee) RegisterEmp() revel.Result {
	reqBody:=c.Request
	err:=reqBody.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	hash,_:=HashPassword(reqBody.FormValue("password"))
	emp := models.EmployeeData{
		FirstName:      reqBody.FormValue("first_name"),
		LastName:       reqBody.FormValue("last_name"),
		Email:          reqBody.FormValue("email"),
		Password:       hash,
		Phone:          reqBody.FormValue("phone"),
		RegistrationAt: time.Now().Unix(),

	}
	empRegister :=models.RegisterEmp(emp)
	if empRegister ==true{
		return c.RenderText("User successfully register ")
	}

	return c.RenderText("Somthing was wrong")
}
