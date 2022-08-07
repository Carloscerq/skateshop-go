package services

import (
    "SkateShop/models"
)

var dbConnection = models.DbConnection

func CreateUser(user *models.User) {
    dbConnection.Create(user)
}
