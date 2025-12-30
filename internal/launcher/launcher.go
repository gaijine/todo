package launcher

import (
	"fmt"
	"log"
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
		fmt.Println("[5] удалить задачу")
		fmt.Println("[0] выйти")
		fmt.Println("-------------------------")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			tasks.AddTask()
		case 2:
			tasks.ShowTasks()
		case 3:
			tasks.MarkDoneTask()
		case 4:
			tasks.UnmarkDoneTask()
		case 5:
			tasks.DeleteTask()
		case 0:
			log.Println("[INFO] application finished")
			return
		default:
			fmt.Println("-------------------------")
			log.Println("[ERROR] invalid menu number")
			fmt.Println("ввели неверный номер")
			fmt.Println("-------------------------")
		}

	}
}
