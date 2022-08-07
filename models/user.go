package models

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    CreditCard string `json:"credit_card"`
    Address string `json:"address"`
    Phone string `json:"phone"`
    Role string `json:"role"`
}
