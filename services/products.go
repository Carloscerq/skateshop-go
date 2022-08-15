package services

import (
	"SkateShop/dto"
	"SkateShop/models"
	"errors"
	"net/http"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func CreateProduct(product *dto.Product) (int, error) {
    productUUID := uuid.New().String()
    newProduct := models.Product {
        UUID:  productUUID,
        Name: product.Name,
        Description: product.Description,
        Image: product.Image,
        Category: product.Category,
        Price: product.Price,
        Stock: 0,
    }
    result := models.DbConnection.Create(&newProduct)
    if result.Error != nil {
        return http.StatusInternalServerError, errors.New("Error creating product")
    }
    return http.StatusOK, nil
}

func GetProduct(uuid string) (*models.Product, error) {
    product := models.Product{}
    models.DbConnection.First(&product, "uuid = ?", uuid)
    if product.Description == "" {
        return nil, errors.New("Product not found")
    }
    return &product, nil
}

func GetProducts() ([]models.Product) {
    var products []models.Product
    models.DbConnection.Find(&products)
    return products
}

func DeleteProduct(uuid string) (error) {
    models.DbConnection.Where("uuid = ?", uuid).Delete(&models.Product{})
    return nil
}

func UpdateProduct(product *dto.Product, uuid string) (error) {
    oldProduct, err := GetProduct(uuid); if err != nil {
        return err
    }
    copier.Copy(&oldProduct, product)
    models.DbConnection.Save(&oldProduct)
    return nil
}

func ChangeStock(request *dto.ProductRequest) (error) {
    product, err := GetProduct(request.ProductID); if err != nil {
        return err
    }
    if uint(request.Amount) + product.Stock < 0 {
        return errors.New("Not enough stock")
    }
    product.Stock += uint(request.Amount)
    models.DbConnection.Save(&product)
    return nil
}
