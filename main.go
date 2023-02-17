package main

import (
	"fmt"
	"os"
	"time"

	infra "go-message-cli/infra/mysql"
	"go-message-cli/model"
	"go-message-cli/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/go_practice?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var f repository.MessageRepository = &infra.MessageRepository{DB: db}

	args := os.Args[1]

	switch args {
	case "new":
		name := os.Args[2]
		message := os.Args[3]
		f.Save(&model.Message{Name: name, Message: message, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	case "list":
		messages, err := f.List()
		if err != nil {
			fmt.Println(err)
		}

		for _, message := range messages {
			fmt.Println(message.Name, message.Message, message.CreatedAt)
		}
	default:
		fmt.Println("Command Not Found")
	}
}
