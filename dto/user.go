package dto

type UserRole string
const (
    USER UserRole = "user"
    ADMIN UserRole = "admin"
)

type NewUser struct {
    Username string
    Email string
    Password string
    CreditCard string
    Address string
    Phone string
    Role UserRole
}

type UpdateUser struct {
    Username string
    Email string
    Password string
    CreditCard string
    Address string
    Phone string
}
