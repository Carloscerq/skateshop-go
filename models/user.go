package models

import (
    "SkateShop/dto"
)

type User struct {
    UUID string `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
    CreditCard string `json:"credit_card"`
    Address string `json:"address"`
    Phone string `json:"phone" gorm:"unique"`
    Role dto.UserRole `json:"role"`
}
