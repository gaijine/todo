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
		fmt.Println("1. создать задачу")
		fmt.Println("2. вывести список задач")
		fmt.Println("3. пометить выполненной")
		fmt.Println("4. снять отметку")
		fmt.Println("0. выйти")
		fmt.Println("-------------------------")
		fmt.Scanln(&choice)

		var text string

		switch choice {

		case 1:
			fmt.Println("-------------------------")
			fmt.Println("введите текст задачи:")
			fmt.Println("-------------------------")
			fmt.Scanln(&text)
			task := tasks.Task{Text: text, Done: false}
			tasks.List = append(tasks.List, task)
			log.Printf("[INFO] task created: %s", task.Text)
		case 2:
			tasks.ShowTasks()
		case 3:
			tasks.MarkDoneTask()
		case 4:
			tasks.UnmarkDoneTask()
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
