package dto

import (
    "github.com/golang-jwt/jwt"
)

type UserClaim struct {
    jwt.StandardClaims
    UserId string `json:"id"`
    UserEmail string `json:"email"`
    UserName string `json:"username"`
}

type UserLoginDto struct {
    Email string `json:"email"`
    Password string `json:"password"`
}
