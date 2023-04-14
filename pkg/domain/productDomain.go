package domain

type Product struct {
	Id          string  `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string  `json:"name" gorm:"not null;unique" binding:"required"`
	Description string  `json:"description"  binding:"required,min=15"`
	Price       float64 `json:"price" gorm:"not null;unique" binding:"required"`
	Stock       int     `json:"stock" gorm:"not null;unique" binding:"required"`
}
