package infra

import (
	"encoding/json"
	"fmt"
	"os"

	"go-message-cli/model"
)

type MessageRepository struct {
	FilePath string
}

// 書き込み
func (m *MessageRepository) Save(msg *model.Message) error {
	f, err := os.OpenFile(m.FilePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	j, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = fmt.Fprint(f, string(j))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// 読み込み
func (m *MessageRepository) List() ([]*model.Message, error) {
	f, err := os.Open(m.FilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var msg []*model.Message
	decoder := json.NewDecoder(f)

	for decoder.More() {
		var message model.Message
		err = decoder.Decode(&message)
		if err != nil {
			return nil, err
		}
		msg = append(msg, &message)
	}

	return msg, err
}
