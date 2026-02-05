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
		fmt.Println("Задача выолнена!")
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("[WARNING]: задайте id задачи и укажите изменение")
			return
		}

		id := os.Args[2]
		taskName := strings.Join(os.Args[3:], " ")
		err := prefixes.TaskUpdate(id, taskName)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			return
		}
		fmt.Println("Задача изменена!")
	default:
		fmt.Println("[ERROR]: неизвестная команда", command)
	}
}
