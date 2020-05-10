package model

type Category struct {
	CategoryId   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type CategoryRepository interface {
	Store(category Category) (Category, error)
	FindById(id int64) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	Delete(category Category) (bool, error)
}
