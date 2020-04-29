package domain

type Client struct {
	UserId      int64
	Username    string
	Password    string
	Description string
}

type ClientRepository interface {
	Store(client Client) Client
	FindById(id int64) Client
}
