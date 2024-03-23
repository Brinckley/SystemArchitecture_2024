package internal

type CreatePostRequest struct {
	AccountId int    `json:"account_id"`
	Content   string `json:"content"`
}

type Post struct {
	Id        int    `json:"id"`
	AccountId int    `json:"account_id"`
	Content   string `json:"content"`
}

func PostFrom(id int, createPost *CreatePostRequest) *Post {
	return &Post{
		Id:        id,
		AccountId: createPost.AccountId,
		Content:   createPost.Content,
	}
}
