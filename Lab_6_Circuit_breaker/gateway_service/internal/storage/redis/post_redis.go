package cache

import (
	"Gateway/internal/entity"
	"Gateway/internal/storage"
	"fmt"
)

const (
	POST_REDIS_PREFIX               = "postId:"
	POSTS_WITH_ACCOUNT_REDIS_PREFIX = "accountWithPosts:"
)

func (r *RedisCache) GetPost(key string) (post entity.Post, err error) {
	err = r.client.Get(r.ctx, POST_REDIS_PREFIX+key).Scan(&post)
	if err != nil {
		return post, storage.NewCacheError(fmt.Sprintf("cannot find post error : %s", err))
	}
	return post, nil
}

func (r *RedisCache) SetPost(key string, post entity.Post) error {
	statusCmd := r.client.Set(r.ctx, POST_REDIS_PREFIX+key, post, r.ttl)
	return statusCmd.Err()
}

func (r *RedisCache) GetPostCollection(accountId string) (posts entity.PostCollection, err error) {
	err = r.client.Get(r.ctx, POSTS_WITH_ACCOUNT_REDIS_PREFIX+accountId).Scan(&posts)
	if err != nil {
		return posts, storage.NewCacheError(fmt.Sprintf("cannot find posts for accountId %s error : %s", accountId, err))
	}
	return posts, nil
}

func (r *RedisCache) SetPostCollection(accountId string, posts entity.PostCollection) error {
	statusCmd := r.client.Set(r.ctx, POSTS_WITH_ACCOUNT_REDIS_PREFIX+accountId, posts, r.ttl)
	return statusCmd.Err()
}
