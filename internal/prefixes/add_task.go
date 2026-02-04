package prefixes

import (
	"encoding/json"
	"fmt"
	"os"
	"task_cli/internal/model"
	"time"
)

// ======= Добавить задачу =======
func Write(taskName string) error {
	var tasks []model.Task
	//читаем tasks.json
	data, err := os.ReadFile("storage/tasks.json")
	if err != nil {
		fmt.Println("[ERROR]: ошибка чтения файла")
		return err
	}
	json.Unmarshal(data, &tasks)

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	newTask := model.Task{
		ID:        id,
		Name:      taskName,
		Done:      false,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, newTask)

	data, err = json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/tasks.json", data, 0644)
}
