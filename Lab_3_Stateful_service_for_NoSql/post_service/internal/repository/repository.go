package repository

import "post_service/internal"

type Storage interface {
	CreatePost(*internal.CreatePostRequest) (int, error)
	GetPostsByAccountId(int) ([]internal.Post, error)
	GetPostByAccountById(int, int) (*internal.Post, error)
	DeletePostByAccountById(int, int) (int, error)
	UpdatePostByAccountById(int, int, string) (*internal.Post, error)
}
