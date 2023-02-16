package infra

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"go-message-cli/model"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func (m *MessageRepository) Save(msg *model.Message) error {
	message := &model.Message{
		Name:      msg.Name,
		Message:   msg.Message,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
	}

	_, err := m.DB.NamedExec(`
		INSERT INTO messages(name, message, created_at, updated_at) 
		VALUES(:name, :message, :created_at, :updated_at)
		`, &message)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (m *MessageRepository) List() ([]*model.Message, error) {
	messages := []*model.Message{}
	err := m.DB.Select(&messages, "SELECT * FROM messages")
	if err != nil {
		fmt.Println(err)
	}

	return messages, nil
}
