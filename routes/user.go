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
    router.DELETE("/:email", deleteUser)
    router.POST("/update/:id", updateUser)
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

func getUser(c *gin.Context) {
    email := c.Param("email")
    user, err := services.GetUser(email); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := services.UpdateUser(&user, userId); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
