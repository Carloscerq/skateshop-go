package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "SkateShop/services"
    "SkateShop/dto"
)

func UserRouterGroup(router *gin.RouterGroup) {
    router.POST("/", createUser)
    router.GET("/:email", getUser)
    router.DELETE("/:email", services.LoginMiddleware(), deleteUser)
    router.PATCH("/:id", services.LoginMiddleware(), updateUser)
}

func createUser(c *gin.Context) {
    var newUser dto.NewUser
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    response, err := services.CreateUser(&newUser); if err != nil {
        c.JSON(response, gin.H{"message": err.Error()})
        return
    }
    c.JSON(response, gin.H{"message": "User created successfully"})
}

func getUser(c *gin.Context) {
    email := c.Param("email")
    user, err := services.GetUser(email); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
    email := c.Param("email")
    err := services.DeleteUser(email); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func updateUser(c *gin.Context) {
    var user dto.UpdateUser
    userId := c.Param("id")
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    err := services.UpdateUser(&user, userId); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
