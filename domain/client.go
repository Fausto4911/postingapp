package domain

type Client struct {
	UserId      int64
	Username    string
	Password    string
	Description string
}

type ClientRepository interface {
	Store(client Client) (Client, error)
	FindById(id int64) (Client, error)
	FindAll() ([]Client, error)
	Update(client Client) (Client, error)
	Delete(client Client) (bool, error)
}
