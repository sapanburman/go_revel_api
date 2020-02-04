package models

import (
	"fmt"
	"revelapi/app/dbconfig"
)


type EmployeeData struct {

	FirstName      string	`json:"first_name"`
	LastName       string	`json:"last_name"`
	Email          string	`json:"email"`
	Password       string	`json:"password"`
	Phone          string	`json:"phone"`
	RegistrationAt int64	`json:"registration_at"`
	UpdateAt       int64	`json:"update_at"`
}

func RegisterEmp(empdata EmployeeData )bool{
_, err := dbconfig.DB.Exec(`INSERT INTO Employee(first_name,last_name, email,password,phone,registration_at) VALUES(?,?,?,?,?,?)`,empdata.FirstName,empdata.LastName,empdata.Email,empdata.Password,empdata.Phone,empdata.RegistrationAt)
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	return true
}
