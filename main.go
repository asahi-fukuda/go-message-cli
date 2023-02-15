package main

import (
	"os"

	infra "go-message-cli/infra/localfile"
)

func main() {
	f := infra.MessageRepository{FilePath: "tmp/message.txt"}

	args := os.Args[1]

	if args == "new" {
		name := os.Args[2]
		message := os.Args[3]
		f.Save(name, message)
	}

	if args == "list" {
		f.List()
	}
}
