package main

import (
	"fmt"
	"os"
	"strings"

	"task_cli/internal/action"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[WARNING]: слишком мало аргументов(add, delete, update)\n[EXAMPLE]: go run main.go (prefix) (task)")
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
		err := action.Write(taskName)
		if err != nil {
			fmt.Println("[ERROR]:", err)
			return
		}
		fmt.Println("Задача добавлена:", taskName)
	default:
		fmt.Println("[ERROR]: неизвестная команда", command)
	}
}
