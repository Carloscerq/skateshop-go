package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ProductRouterGroup(router *gin.RouterGroup) {
    router.GET("/", GetProducts)
    router.GET("/:id", GetProduct)
    router.POST("/", CreateProduct)
    router.PATCH("/:id", UpdateProduct)
    router.DELETE("/:id", DeleteProduct)
}

func GetProducts(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"message": "GetProducts"})
}

func GetProduct(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"message": "GetProduct"})
}

func CreateProduct(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"message": "CreateProduct"})
}

func UpdateProduct(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"message": "UpdateProduct"})
}

func DeleteProduct(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"message": "DeleteProduct"})
}
