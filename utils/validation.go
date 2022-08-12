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

var PriceValidator validator.Func = func(fl validator.FieldLevel) bool {
    data, ok := fl.Field().Interface().(float64)
    if !ok {
        return false
    }
    return data >= 1.0
}

var categoryRegex = regexp.MustCompile(`(shape|wheel|hardware|griptape)`)
var CategoryValidator validator.Func = func(fl validator.FieldLevel) bool {
    data, ok := fl.Field().Interface().(dto.Category)
    if !ok {
        return false
    }
    return categoryRegex.MatchString(string(data))
}
