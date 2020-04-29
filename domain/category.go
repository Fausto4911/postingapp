package domain

type Category struct {
	CategoryId   int64
	CategoryName string
}

type CategoryRepository interface {
	Store(category Category) Category
	FindById(id int64) Category
}
