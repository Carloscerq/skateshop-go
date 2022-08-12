package models

type Product struct {
    UIID string `json:"id" gorm:"primary_key"`
    Name string `json:"name"`
    Description string `json:"description"`
    Price float64 `json:"price"`
    Image string `json:"image"`
    Category string `json:"category"`
}
