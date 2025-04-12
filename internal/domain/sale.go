package domain

type Sale struct {
	Model
	ProductName string  `json:"product_name" gorm:"not null"`
	Quantity    int     `json:"quantity" gorm:"default:0 not null unsigned"`
	Price       float64 `json:"price" gorm:"not null"`
	Total       float64 `json:"total" gorm:"not null"`
	User        *User   `json:"user" gorm:"foreignKey:UserID"`
}

func NewSale(productName string, quantity int, price float64) *Sale {
	return &Sale{
		ProductName: productName,
		Quantity:    quantity,
		Price:       price,
		Total:       float64(quantity) * price,
	}
}
