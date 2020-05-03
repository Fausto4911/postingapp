package domain

type Category struct {
	CategoryId   int64
	CategoryName string
}

type CategoryRepository interface {
	Store(category Category) (Category, error)
	FindById(id int64) (Category, error)
	FindAll() ([]Category, error)
	Update(category Category) (Category, error)
	Delete(category Category) (bool, error)
}
