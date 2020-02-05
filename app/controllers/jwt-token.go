package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"log"
	"revelapi/app/models"
)


type JWT_TOKEN struct {
	*revel.Controller
}
func (c JWT_TOKEN) token() revel.Result {
	reqBody:=c.Request
	err:=reqBody.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	eamil:=reqBody.FormValue("email")
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	tokenReq.RefreshToken=reqBody.FormValue("refresh_token")

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["isDisabled"].(float64)) == 0 {

			newTokenPair, err := models.GenerateTokenPair(eamil)
			if err != nil {
				return c.RenderText("Somthing was wrong")
			}

			return c.RenderJSON( newTokenPair)
		}

		return c.RenderText("ErrUnauthorized")
	}

	return c.RenderText("check")
}