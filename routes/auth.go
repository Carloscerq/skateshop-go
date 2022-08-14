package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "SkateShop/services"
    "SkateShop/dto"
)

func AuthRouterGroup(router *gin.RouterGroup) {
    router.POST("/login", login)
}

func login(c *gin.Context) {
    var userLoginDto dto.UserLoginDto
    if err := c.ShouldBindJSON(&userLoginDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    token, err := services.Login(userLoginDto.Email, userLoginDto.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
