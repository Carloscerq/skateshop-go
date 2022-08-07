package models

import (
    "gorm.io/gorm"
    "SkateShop/dto"
)

type User struct {
    gorm.Model
    Username string `json:"username"`
    Email    string `json:"email" gorm:"unique_index"`
    Password string `json:"password"`
    CreditCard string `json:"credit_card"`
    Address string `json:"address"`
    Phone string `json:"phone" gorm:"unique_index"`
    Role dto.UserRole `json:"role"`
}
