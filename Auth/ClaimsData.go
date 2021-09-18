package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("jwt_private_key")

type JWT struct {
	StatusCode int    `json:"statuscode"`
	Token      string `json:"token"`
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email      string
	Role       string
	User_ID    int
	Company_id int
	jwt.StandardClaims
}
