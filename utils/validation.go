package utils

import (
    "SkateShop/dto"
    "regexp"
    "github.com/go-playground/validator/v10"
)

var roleRegex = regexp.MustCompile(`(user|admin)`)
var RoleValidator validator.Func = func(fl validator.FieldLevel) bool {
    data, ok := fl.Field().Interface().(dto.UserRole)
    if !ok {
        return false
    }
    return roleRegex.MatchString(string(data))
}
