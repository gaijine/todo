package launcher

import (
	"fmt"
	"log"
	"todo/internal/storage"
	"todo/internal/tasks"
)

func Menu() {
	var choice int
	for {
		fmt.Println("-------------------------")
		fmt.Println("[1] создать задачу")
		fmt.Println("[2] вывести список задач")
		fmt.Println("[3] пометить выполненной")
		fmt.Println("[4] снять отметку")
		fmt.Println("[5] изменить текст задачи")
		fmt.Println("[6] удалить задачу")
		fmt.Println("[0] выйти")
		fmt.Println("-------------------------")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			tasks.AddTask()
			storage.SaveTasksToFile()
		case 2:
			tasks.ShowTasksMenu()
		case 3:
			tasks.MarkDoneTask()
			storage.SaveTasksToFile()
		case 4:
			tasks.UnmarkDoneTask()
			storage.SaveTasksToFile()
		case 5:
			tasks.UpdateTask()
			storage.SaveTasksToFile()
		case 6:
			tasks.DeleteTask()
			storage.SaveTasksToFile()
		case 0:
			log.Println("[INFO] application finished")
			return
		default:
			fmt.Println("-------------------------")
			log.Println("[ERROR] invalid menu number")
			fmt.Println("выбран неверный номер")
			fmt.Println("-------------------------")
		}

	}
}
