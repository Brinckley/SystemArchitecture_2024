package cache

import "Gateway/internal/entity"

type SNCache interface {
	GetAccount(key string) (entity.Account, error)
	SetAccount(key string, value entity.Account) error
	DeleteAccount(key string) error

	GetMessage(key string) (entity.Message, error)
	SetMessage(key string, value entity.Message) error
	GetMessageCollection(key string) (entity.MessagesCollection, error)
	SetMessageCollection(key string, value entity.MessagesCollection) error

	GetPost(key string) (entity.Post, error)
	SetPost(key string, value entity.Post) error
	DeletePost(key string) error
	GetPostCollection(key string) (entity.PostCollection, error)
	SetPostCollection(key string, value entity.PostCollection) error
}
