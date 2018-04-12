package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"log"
)

func main() {

	// ********* SENDER **********

	// Token - JWT Header
	hmac512Algo := jwt.SigningMethodHS512.Alg()
	signingMethod := jwt.GetSigningMethod(hmac512Algo)
	token := jwt.New(signingMethod)

	// Claims  - JWT Payload
	claims := make(jwt.MapClaims)
	claims["test"] = "this is the test data"
	token.Claims = claims

	// 256 bit secret key for HMAC512
	secretKey := securecookie.GenerateRandomKey(256)

	// Signing and Serialization.
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Print("Error: ", err)
		return
	}

	log.Printf("token string: %+v", tokenString)

	//******* RECEIVER *******

	// Token received after Decrypting from token string and secret key.
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("error: %+v", err)
		return
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		log.Printf("Claims: %+v", claims)
	} else {
		fmt.Println(err)
	}

}
