package services

import (
    "github.com/google/uuid"
    "github.com/jinzhu/copier"
    "SkateShop/models"
    "SkateShop/dto"
    "SkateShop/utils"
    "net/http"
    "errors"
)

func CreateUser(user *dto.NewUser) (int, error) {
    hashPassword, err := utils.Hash(user.Password)
    userUUID := uuid.New().String()
    if err != nil {
        return http.StatusInternalServerError, err
    }
    newUser := models.User{
        UUID: userUUID,
        Username: user.Username,
        Password: hashPassword,
        Email: user.Email,
        CreditCard: user.CreditCard,
        Address: user.Address,
        Phone: user.Phone,
        Role: user.Role,
    }
    result := models.DbConnection.Create(&newUser)
    if result.Error != nil {
        return http.StatusInternalServerError, errors.New("Error creating user")
    }
    return http.StatusOK, nil
}

func GetUser(email string) (*models.User, error) {
    user := models.User{}
    models.DbConnection.First(&user, "email = ?", email)
    if user.Email == "" {
        return nil, errors.New("User not found")
    }
    return &user, nil
}

func GetUserByID(id string) (*models.User, error) {
    user := models.User{}
    models.DbConnection.First(&user, "uuid = ?", id)
    if user.Email == "" {
        return nil, errors.New("User not found")
    }
    return &user, nil
}

func DeleteUser(email string) (error) {
    models.DbConnection.Where("email = ?", email).Delete(&models.User{})
    return nil
}

func UpdateUser(user *dto.UpdateUser, id string) (error) {
    oldUser, err := GetUserByID(id); if err != nil {
        return err
    }
    copier.Copy(&oldUser, user)
    models.DbConnection.Save(&oldUser)
    return nil
}

func addToken(token string, userid string) error {
    user, err := GetUserByID(userid); if err != nil {
        return err
    }
    user.Token = token
    models.DbConnection.Save(&user)
    return nil
}
