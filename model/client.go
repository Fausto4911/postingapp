package model

type Client struct {
	UserId      int64  `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type ClientRepository interface {
	Store(client Client) (Client, error)
	FindById(id int64) (Client, error)
	FindAll() ([]Client, error)
	Update(client Client) (Client, error)
	Delete(client Client) (bool, error)
}
