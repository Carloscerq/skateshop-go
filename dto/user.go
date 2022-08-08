package dto

type UserRole string
const (
    USER UserRole = "user"
    ADMIN UserRole = "admin"
)

type NewUser struct {
    Username string `binding:"required" json:"username"`
    Email string `binding:"required" json:"email"`
    Password string `binding:"required" json:"password"`
    CreditCard string `binding:"required" json:"creditCard"`
    Address string `binding:"required" json:"address"`
    Phone string `binding:"required" json:"phone"`
    Role UserRole `binding:"required" json:"role"`
}

type UpdateUser struct {
    Username string `binding:"required" json:"username"`
    Email string `binding:"required" json:"email"`
    Password string `binding:"required" json:"password"`
    CreditCard string `binding:"required" json:"creditCard"`
    Address string `binding:"required" json:"address"`
    Phone string `binding:"required" json:"phone"`
}
