package main

import (
	"fmt"
	"log"

	"github.com/gorilla/securecookie"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	// ======== Sender ===========
	message := make(map[string]interface{})
	message["World"] = "Hello"
	secretKey := securecookie.GenerateRandomKey(256)
	token := EncJWT256(message, secretKey)

	// ======== Reciver ===========
	claims, err := DecJWT256(token, secretKey)
	if err != nil {
		fmt.Printf("\n err :: %+v \n\n", err)
	}

	fmt.Printf("\n claims :: %+v \n\n", claims)

}

// EncJWT256 encrypts the message with the given secret key and returns the 'Json Web Token (JWT)'
// secretkey lenght depends on the type of encryption algoritham. (i.e for HMAC512 the secret key length is 512 byte)
// Example
// 		message := make(map[string]interface{})
//		message["World"] = "Hello"
//		secretKey := securecookie.GenerateRandomKey(256)
//		token := EncJWT256(message, secretKey)
func EncJWT256(message map[string]interface{}, secretKey []byte) (JWTToken string) {

	// Token - JWT Header
	hmac512Algo := jwt.SigningMethodHS512.Alg()
	signingMethod := jwt.GetSigningMethod(hmac512Algo)
	token := jwt.New(signingMethod)

	// Claims  - JWT Payload
	claims := make(jwt.MapClaims)
	for key, val := range message {
		claims[key] = val
	}

	token.Claims = claims

	// Signing and Serialization.
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Print("Error: ", err)
		return
	}

	return tokenString
}

// DecJWT256 decrypts the Json Web Token using the secret key and returns the message
// secretKey which given for encryption is must be given in params
// Example
//		claims, err := DecJWT256(token, secretKey)
//		if err != nil {
//			fmt.Printf("\n err :: %+v \n\n", err)
//		}
//
//		fmt.Printf("\n claims :: %+v \n\n", claims)
func DecJWT256(tokenString string, secretKey []byte) (claims jwt.Claims, err error) {

	// Token received after Decrypting from token string and secret key.
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if parsedToken.Valid {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
			return claims, nil
		}
	}

	return nil, fmt.Errorf("Error: %s \n", "failed to parse tokenString")
}
