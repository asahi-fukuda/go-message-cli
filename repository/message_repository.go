package repository

import "go-message-cli/model"

type MessageRepository interface {
	// Message型
	Save(message *model.Message) error
	// Message型の配列
	List() (*[]model.Message, error)
}
