package services

import (
	"SkateShop/dto"
	"SkateShop/models"
	"SkateShop/utils"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func LoginMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := strings.Split(c.Request.Header.Get("Authorization"), " ")
        if len(token) != 2 {
            c.AbortWithStatus(401)
            return
        }
        claims := &dto.UserClaim{}
        _, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        } )
        if err != nil {
            c.AbortWithStatus(401)
            return
        }
        c.Set("user", claims)
    }
}
