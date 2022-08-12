package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func AuthRouterGroup(router *gin.RouterGroup) {
    router.POST("/login", login)
}

func login(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}
