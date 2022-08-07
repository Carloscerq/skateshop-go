package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "SkateShop/services"
    "SkateShop/dto"
)

func UserRouterGroup(router *gin.RouterGroup) {
    router.POST("/", createUser)
}

func createUser(c *gin.Context) {
    var newUser dto.NewUser
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    _, err := services.CreateUser(&newUser); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
