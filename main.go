package main

import (
	"fmt"
	"os"
	"strings"

	"task_cli/internal/prefixes"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[WARNING]: слишком мало аргументов(add, done, delete, update)\n[EXAMPLE]: go run main.go (prefix) (task)")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("[WARNING]: добавьте задание")
			return
		}
		taskName := strings.Join(os.Args[2:], " ")
		err := prefixes.Write(taskName)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			return
		}
		fmt.Println("Задача добавлена:", taskName)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("[WARNING]: укажите id задания")
			return
		}
		id := os.Args[2]
		err := prefixes.TaskDoneUpdate(id)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			return
		}
	default:
		fmt.Println("[ERROR]: неизвестная команда", command)
	}
}
