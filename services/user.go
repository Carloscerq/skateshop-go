package services

import (
    "SkateShop/models"
    "SkateShop/dto"
    "SkateShop/utils"
    "net/http"
)

func CreateUser(user *dto.NewUser) (int, error) {
    hashPassword, err := utils.Hash(user.Password)
    if err != nil {
        return http.StatusInternalServerError, err
    }
    user.Password = hashPassword
    models.DbConnection.Create(&models.User{
        Username: user.Name,
        Password: user.Password,
        Email: user.Email,
        CreditCard: user.CreditCard,
        Address: user.Address,
        Phone: user.Phone,
        Role: user.Role,
    })
    return http.StatusOK, nil
}
