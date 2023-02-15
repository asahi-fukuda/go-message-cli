package main

import (
	"fmt"
	"os"
	"time"

	infra "go-message-cli/infra/localfile"
	"go-message-cli/model"
	"go-message-cli/repository"
)

func main() {
	var f repository.MessageRepository = &infra.MessageRepository{FilePath: "tmp/message.txt"}

	args := os.Args[1]

	switch args {
	case "new":
		name := os.Args[2]
		message := os.Args[3]
		f.Save(model.Message{Name: name, Message: message, Time: time.Now()})
	case "list":
		f.List()
	default:
		fmt.Println("Command Not Found")
	}
}
