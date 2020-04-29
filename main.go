package main

import (
	"fmt"
	"postingapp/domain"
)

func main() {
	fmt.Println("======== Running main ==========")
	var cl domain.Client
	cl.UserId = 78
	fmt.Println(cl, cl.Username)

	var post domain.Post
	post.Client = cl
	fmt.Println(post)
}
