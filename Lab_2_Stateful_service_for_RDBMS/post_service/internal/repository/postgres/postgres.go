package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"post_service/internal"
)

type PostgresStorage struct {
	db        *sql.DB
	tableName string
}

func NewPostgresStorage(table string) (*PostgresStorage, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password='%s' sslmode=disable search_path=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SCHEMA"))
	postgresDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := postgresDB.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db:        postgresDB,
		tableName: table,
	}, nil
}

func (p PostgresStorage) CreatePost(createPost *internal.CreatePostRequest) (int, error) {
	queryInsertPost := fmt.Sprintf(
		"insert into %s (account_id, content) values ($1, $2) returning id;", p.tableName)
	log.Println(queryInsertPost)
	var postId int
	err := p.db.QueryRow(queryInsertPost, createPost.AccountId, createPost.Content).Scan(&postId)
	if err != nil {
		return -1, err
	}
	return postId, nil
}

func (p PostgresStorage) GetPostsByAccountId(accountId int) ([]internal.Post, error) {
	selectAllQuery := fmt.Sprintf("SELECT * FROM %s WHERE account_id=%d;", p.tableName, accountId)
	log.Println(selectAllQuery)
	rows, err := p.db.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}

	var posts []internal.Post
	for rows.Next() {
		post := new(internal.Post)
		err := rows.Scan(
			&post.Id,
			&post.AccountId,
			&post.Content,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *post)
	}
	return posts, nil
}

func (p PostgresStorage) GetPostByAccountById(accountId int, postId int) (*internal.Post, error) {
	var post internal.Post
	selectById := fmt.Sprintf("SELECT * FROM %s WHERE id=%d AND account_id=%d;", p.tableName, postId, accountId)
	log.Println(selectById)
	err := p.db.QueryRow(selectById).Scan(
		&post.Id,
		&post.AccountId,
		&post.Content,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p PostgresStorage) UpdatePostByAccountById(accountId, postId int, content string) (*internal.Post, error) {
	updateQuery := fmt.Sprintf("UPDATE %s SET content=$3 WHERE id=$1 AND account_id=$2;", p.tableName)
	log.Println(updateQuery)
	_, err := p.db.Exec(updateQuery, postId, accountId, content)
	if err != nil {
		return nil, err
	}
	return p.GetPostByAccountById(postId, accountId)
}

func (p PostgresStorage) DeletePostByAccountById(accountId int, postId int) (int, error) {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND account_id=$2;", p.tableName)
	_, err := p.db.Exec(deleteQuery, postId, accountId)
	if err != nil {
		return -1, err
	}
	return postId, nil
}
