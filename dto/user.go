package dto

type UserRole string
const (
    USER UserRole = "user"
    ADMIN UserRole = "admin"
)

type NewUser struct {
    Name string
    Email string
    Password string
    CreditCard string
    Address string
    Phone string
    Role UserRole
}
