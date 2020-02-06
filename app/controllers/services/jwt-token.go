package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"log"
	"time"
)


type JwtToken struct {
	*revel.Controller
}

// secret key (note: good practice get from env)
var mySigningKey = []byte{4,5,6,78,8,5,67,9}

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


func (c JwtToken) Token() revel.Result {
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
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["isDisabled"].(float64)) == 0 {

			newTokenPair, err := GenerateTokenPair(eamil)
			if err != nil {
				return c.RenderText("Somthing was wrong")
			}

			return c.RenderJSON( newTokenPair)
		}

		return c.RenderText("ErrUnauthorized")
	}

	return c.RenderText("check")
}