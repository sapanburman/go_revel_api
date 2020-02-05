package models

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"log"
	"revelapi/app/controllers"
	"revelapi/app/dbconfig"
	"time"
)

// secret key (note: good practice get from env)
var mySigningKey = []byte{4,5,6,78,8,5,67,9}

// Create a struct to read the email and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Email string `json:"email"`
}

func GenerateTokenPair(email string) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get user etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["isDisabled"] = 0
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshTokenString, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

func CheckPass(creds Credentials)bool{
	hashpass := &Credentials{}
	// Get the existing entry present in the database for the given email
	err := dbconfig.DB.QueryRow("SELECT password from Employee where email=?", creds.Email).Scan(&hashpass.Password)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	isValid:=controllers.CheckPasswordHash(creds.Password,hashpass.Password)
	if isValid{
		return true
	}
	return false
}