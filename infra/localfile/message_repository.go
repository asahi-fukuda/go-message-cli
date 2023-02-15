package infra

import (
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

	_, err = fmt.Fprintln(f, *msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// 読み込み
func (m *MessageRepository) List() ([]*model.Message, error) {
	const BUFFERSIZE = 255

	f, err := os.Open(m.FilePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	b := make([]byte, BUFFERSIZE)
	for {
		n, err := f.Read(b)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	fmt.Println(string(b))

	return nil, err
}
