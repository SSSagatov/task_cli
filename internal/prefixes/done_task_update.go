package prefixes

import (
	"encoding/json"
	"fmt"
	"os"
	"task_cli/internal/model"
)

// ======= Выполнено/Не выполнено =======
func TaskDoneUpdate(id string) error {
	var tasks *model.Task
	//читаем tasks.json
	data, err := os.ReadFile("storage/tasks.json")
	if err != nil {
		fmt.Println("[ERROR]: ошибка чтения файла при обновлении статуса задания")
		return err
	}
	json.Unmarshal(data, &tasks)

}
