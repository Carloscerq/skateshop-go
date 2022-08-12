package dto

type Category string
const (
    SHAPE Category = "shape"
    WHEEL Category = "wheel"
    HARDWARE Category = "hardware"
    GRIPTAPE Category = "griptape"
)

type Product struct {
    Name string `binding:"required" json:"name"`
    Description string `binding:"required" json:"description"`
    Price float64 `binding:"required,price" json:"price"`
    Image string `json:"image"`
    Category Category `binding:"category" json:"category"`
}

