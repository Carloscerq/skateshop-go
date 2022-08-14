package services

import (
    "SkateShop/dto"
    "SkateShop/models"
    "SkateShop/utils"
    "errors"
    "time"
    "os"
    "github.com/golang-jwt/jwt"
)

func Login(email string, password string) (string, error) {
    user, err := GetUser(email)
    if err != nil || !utils.CheckHash(password, user.Password) {
        return "", errors.New("Invalid credentials")
    }

    token, err := generateToken(user)
    if err != nil {
        return "", errors.New("Error generating token")
    }
    addToken(token, user.UUID)
    return token, nil
}

func generateToken(user *models.User) (string, error) {
    claims := dto.UserClaim{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            Issuer:    "SkateShop",
        },
        UserId: user.UUID,
        UserEmail: user.Email,
        UserName: user.Username,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
