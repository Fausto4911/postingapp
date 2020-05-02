package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"postingapp/domain"
	"strconv"
)

type PostInteractor interface {
	AddClient(userName, password string) error
	AddPost(clientId int64, header, category string) error
	AddComment(postId, clientId int64, description string) error
	AddVote(postId int64) error
	GetClientPostsById(clientId int64) ([]domain.Post, error)
}

type WebserviceHandler struct {
	PostInteractor PostInteractor
}

func (handler WebserviceHandler) GetClientPostsById(res http.ResponseWriter, req *http.Request) {
	clientId, _ := strconv.ParseInt(req.FormValue("clientId"), 10, 64)
	var postList, _ = handler.PostInteractor.GetClientPostsById(clientId)
	for _, element := range postList {
		io.WriteString(res, fmt.Sprintf("item PostId: %d\n", element.PostId))
		io.WriteString(res, fmt.Sprintf("item Header: %v\n", element.Header))
		io.WriteString(res, fmt.Sprintf("item Category: %f\n", element.Category))
		io.WriteString(res, fmt.Sprintf("item CreationDate: %f\n", element.CreationDate))
	}

}

// Example code
// func (handler WebserviceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
//     userId, _ := strconv.Atoi(req.FormValue("userId"))
//     orderId, _ := strconv.Atoi(req.FormValue("orderId"))
//     items, _ := handler.OrderInteractor.Items(userId, orderId)
//     for _, item := range items {
//         io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
//         io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
//         io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
//     }
// }
