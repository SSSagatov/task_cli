package prefixes

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"task_cli/internal/model"
)

// ======= Удаление задания =======
func RemoveTask(id string) error {
	//переводим id string в integer
	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("[ERROR]: не удалось получить id")
		return err
	}
	//читаем tasks.json
	data, err := os.ReadFile("storage/tasks.json")
	if err != nil {
		fmt.Println("[ERROR]: ошибка чтения файла при обновлении статуса задания")
		return err
	}
	//парсим json
	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("[ERROR]: ошибка парсинга JSON")
		return err
	}
	//находим по id и удаляем из списка
	found := false
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	//логгируем ошибку если не смогли найти ошибку
	if !found {
		return fmt.Errorf("[ERROR]: задача с id %d не найдена", taskID)
	}
	//сохраняем обратно в файл
	data, err = json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	//записываем обновленный результат
	err = os.WriteFile("storage/tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
