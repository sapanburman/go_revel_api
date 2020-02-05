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

func IsEmailExist(email string)bool{
	var getEmail string
	// Execute the query
	err := dbconfig.DB.QueryRow("SELECT email FROM Employee WHERE email = ?", email).Scan(&getEmail)
	if err != nil {
		panic(err.Error())// proper error handling instead of panic in your app
		return false
	}
	if getEmail==email{
		return true
	}
	return false
}

func IsPhoneExist(phone string)bool{
	var getPhone string
	// Execute the query
	err := dbconfig.DB.QueryRow("SELECT phone FROM Employee WHERE phone = ?", phone).Scan(&getPhone)
	if err != nil {
		panic(err.Error())// proper error handling instead of panic in your app
		return false
	}
	if getPhone==phone{
		return true
	}
	return false
}