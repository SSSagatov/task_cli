package main

import (
	"fmt"
	"os"
	"strings"

	"task_cli/internal/prefixes"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[WARNING]: слишком мало аргументов(add, done, update, remove)\n[EXAMPLE]: go run main.go (prefix) (task)")
		return
	}

	command := os.Args[1]

	switch command {
	// prefix "add" - добавление задачи
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
	// prefix "done" - объявление выполненной задачи
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("[WARNING]: укажите id задания(для выполнения)")
			return
		}

		id := os.Args[2]
		err := prefixes.TaskDoneUpdate(id)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			return
		}
		fmt.Println("Задача выолнена!")
	// prefix "update" - обновление условия задачи
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
	// prefix "remove" - удаление задачи
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("[ERROR]: укажите id задания(для удаления)")
			return
		}

		id := os.Args[2]
		err := prefixes.RemoveTask(id)
		if err != nil {
			fmt.Println("[ERROR]:", err)
		}
		fmt.Println("Задача удалена!")
	// default значение, если никакой из prefix не подошел
	default:
		fmt.Println("[ERROR]: неизвестная команда", command)
	}
}
