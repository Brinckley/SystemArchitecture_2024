package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"message_service/internal"
	"os"
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

func (p *PostgresStorage) CreateMessage(message *internal.CreateMessageRequest) (int, error) {
	queryInsertMessage := fmt.Sprintf(
		"insert into %s (sender_id, receiver_id, content) values ($1, $2, $3) returning id;", p.tableName)
	log.Println(queryInsertMessage)
	var messageId int
	err := p.db.QueryRow(queryInsertMessage, message.SenderId, message.ReceiverId, message.Content).Scan(&messageId)
	if err != nil {
		return -1, err
	}
	return messageId, nil
}

func (p *PostgresStorage) GetMessagesByReceiverId(receiverId int) ([]internal.Message, error) {
	queryGetMessagesByReceiverId := fmt.Sprintf(
		"SELECT * FROM %s WHERE receiver_id=%d;", p.tableName, receiverId)

	rows, err := p.db.Query(queryGetMessagesByReceiverId)
	if err != nil {
		return nil, err
	}

	var messages []internal.Message
	for rows.Next() {
		message := new(internal.Message)
		err := rows.Scan(
			&message.Id,
			&message.SenderId,
			&message.ReceiverId,
			&message.Content,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, *message)
	}
	return messages, nil
}

func (p *PostgresStorage) GetMessageById(receiverId int, messageId int) (*internal.Message, error) {
	queryGetMessageById := fmt.Sprintf(
		"SELECT * FROM %s WHERE receiver_id=%d AND id=%d;", p.tableName, receiverId, messageId)
	var message internal.Message
	err := p.db.QueryRow(queryGetMessageById).Scan(
		&message.Id, &message.SenderId, &message.ReceiverId, &message.Content)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
