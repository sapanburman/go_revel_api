package controllers

import (
	"github.com/revel/revel"
	"log"
	"revelapi/app/controllers/services"
	"revelapi/app/controllers/utils"
	"revelapi/app/models"
	"time"
)


type Employee struct {
	*revel.Controller
}

type Login struct {
	*revel.Controller
}
type Logout struct {
	*revel.Controller
}
func (c Employee) RegisterEmp() revel.Result {
	reqBody:=c.Request
	err:=reqBody.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	hash,_:= services.HashPassword(reqBody.FormValue("password"))
	emp := models.EmployeeData{
		FirstName:      reqBody.FormValue("first_name"),
		LastName:       reqBody.FormValue("last_name"),
		Email:          reqBody.FormValue("email"),
		Password:       hash,
		Phone:          reqBody.FormValue("phone"),
		RegistrationAt: time.Now().Unix(),

	}
	emailCheck:=models.IsEmailExist(emp.Email)
	phoneCheck:=models.IsPhoneExist(emp.Phone)
	if emailCheck{
		return c.RenderJSON(utils.ErrorRes{
			Status:  "error",
			Code:    401,
			Message: "Email already used",
		})
	}
	if phoneCheck{
		return  c.RenderJSON(utils.ErrorRes{
			Status:  "error",
			Code:    401,
			Message: "Phone already used",
		})
	}
	empRegister :=models.RegisterEmp(emp)
	if empRegister ==true{
		return c.RenderJSON(utils.ErrorRes{
			Status:  "success",
			Code:    200,
			Message: "User successfully register ",
		})
	}
	 return c.RenderJSON(utils.ErrorRes{
		Status:  "error",
		Code:    401,
		Message: "Somthing was wrong ",
	})
}

func (c Login) LoginEmp() revel.Result {
	reqBody:=c.Request
	formData:=models.Credentials{
		Password: reqBody.FormValue("password"),
		Email:    reqBody.FormValue("email"),
	}
	c.Session["email"] = formData.Email
	// Check in your db if the user exists or not
	getHash:=models.GetHashPass(formData)
	if getHash !=""{
		isValid:= services.CheckPasswordHash(formData.Password,getHash)
		if isValid{
			tokens, err := services.GenerateTokenPair(formData.Email)
			if err != nil {
				return c.RenderText("Somthing was wrong")
			}

			return c.RenderJSON(utils.SuccessRes{
				Status:  "success",
				Code:    200,
				Data:    tokens,
				Message: "User Successfully login",
			})
		}
		c.RenderJSON(utils.ErrorRes{
			Status:  "error",
			Code:    401,
			Message: "Email or Password Not Found",
		})
	}

	return c.RenderJSON(utils.ErrorRes{
		Status:  "error",
		Code:    401,
		Message: "Email or Password Not Found",
	})
}

func (c Logout) LogoutEmp() revel.Result {

	token := c.Request.GetHttpHeader("Authorization")
	if token == "" {
		log.Fatal("Authorization token was not provided")

		return c.RenderText("Authorization Token is required")

	}

	delete(c.Session,"token")


	return c.RenderText( "Done")


}