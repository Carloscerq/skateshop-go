package models

import (
    "SkateShop/dto"
)

type Product struct {
    UUID string `json:"id" gorm:"primary_key"`
    Name string `json:"name"`
    Description string `json:"description"`
    Price float64 `json:"price"`
    Image string `json:"image"`
    Category dto.Category `json:"category"`
    Stock uint `json:"stock"`
}
