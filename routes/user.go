package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func UserRouterGroup(router *gin.RouterGroup) {
    router.POST("/", createUser)
}

func createUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}
