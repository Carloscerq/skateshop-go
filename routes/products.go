package routes

import (
    "SkateShop/dto"
    "SkateShop/services"
    "net/http"
    "github.com/gin-gonic/gin"
)

func ProductRouterGroup(router *gin.RouterGroup) {
    router.GET("/", getProducts)
    router.GET("/:id", getProduct)
    router.POST("/", services.LoginMiddleware(), createProduct)
    router.PATCH("/:id", services.LoginMiddleware(), updateProduct)
    router.DELETE("/:id", services.LoginMiddleware(), deleteProduct)
    router.POST("/stock", services.LoginMiddleware(), changeStock)
}

func createProduct(c *gin.Context) {
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

func getProduct(c *gin.Context) {
    uuid := c.Param("id")
    if uuid == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
        return
    }
    response, err := services.GetProduct(uuid); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": response})
}

func getProducts(c *gin.Context) {
    response := services.GetProducts()
    c.JSON(http.StatusOK, gin.H{"message": response})
}

func updateProduct(c *gin.Context) {
    var product dto.Product
    uuid := c.Param("id")
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    err := services.UpdateProduct(&product, uuid); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func deleteProduct(c *gin.Context) {
    uuid := c.Param("id")
    if uuid == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UUID"})
        return
    }
    err := services.DeleteProduct(uuid); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func changeStock(c *gin.Context) {
    var productRequest dto.ProductRequest
    if err := c.ShouldBindJSON(&productRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    if productRequest.Amount < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid amount"})
        return
    }

    err := services.ChangeStock(&productRequest); if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Stock updated successfully"})
}
