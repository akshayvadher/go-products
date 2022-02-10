package products

type Product struct {
	Id          int64   `json:"id" gorm:"primary_key"`
	Name        string  `json:"name" validate:"required,min=3,max=50"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
