package infra

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"go-message-cli/model"
)

type MessageRepository struct {
	FilePath string
}

// 書き込み
func (m *MessageRepository) Save(name string, message string) error {
	f, err := os.OpenFile(m.FilePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	msg := model.Message{
		Name:    name,
		Message: message,
		Time:    time.Now(),
	}

	_, err = fmt.Fprintln(f, msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// 読み込み
func (m *MessageRepository) List() error {
	f, err := ioutil.ReadFile(m.FilePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(f))

	return nil
}
