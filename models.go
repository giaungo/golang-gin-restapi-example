package main

type comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type postDto struct {
	PostID           int    `json:"post_id"`
	PostTitle        string `json:"post_title"`
	PostBody         string `json:"post_body"`
	NumberOfComments int    `json:"total_number_of_comments"`
}
