package routes

import (
	"SkateShop/dto"
	"SkateShop/services"
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
    var newProduct dto.Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    response, err := services.CreateProduct(&newProduct); if err != nil {
        c.JSON(response, gin.H{"message": err.Error()})
        return
    }
    c.JSON(response, gin.H{"message": "Product created successfully"})
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
