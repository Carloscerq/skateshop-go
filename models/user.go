package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    CreditCard string `json:"credit_card"`
    Address string `json:"address"`
    Phone string `json:"phone"`
    Role string `json:"role"`
}
