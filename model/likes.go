package model

type Likes struct {
	Id         uint16  `json:"id"`
	User       AppUser `json:"post"`
	TypeId     uint8   `json: "type"` // 0 if like belongs to post and 1 to comment
	Quantity   int32   `json:"quantity"`
	EntityType uint16  `json:"entity"` // the entity id type (could be post id or comment id)
}

type LikesRepository interface {
	Store(likes Likes) error
	GetAll() ([]Likes, error)
	GetById(id int) (Likes, error)
	Delete(id int) error
	Update(likes Likes) error
}
