package tasks

import (
	"fmt"
	"log"
)

func AddTask() {
	var text string
	fmt.Println("-------------------------")
	fmt.Println("введите текст задачи:")
	fmt.Println("-------------------------")
	fmt.Scanln(&text)
	task := Task{Text: text, Done: false}
	List = append(List, task)
	log.Printf("[INFO] task created: %s", task.Text)
}

func DeleteTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	var choice int
	ShowTasks()
	fmt.Println("-------------------------")
	fmt.Println("выберите какую задачу нужно удалить")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(List) {
		deletedTask := List[choice-1]
		List = append(List[0:choice-1], List[choice:]...)
		fmt.Println("-------------------------")
		log.Printf("[INFO] task deleted: id=%d, text=%s", choice, deletedTask.Text)
		fmt.Println("задача успешно удалена")
		fmt.Println("-------------------------")
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран неверный номер")
		fmt.Println("-------------------------")
	}
}

func MarkDoneTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	var choice int
	ShowTasks()
	fmt.Println("-------------------------")
	fmt.Println("выберите какую задачу пометить выполненной")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(List) {
		if List[choice-1].Done != true {
			List[choice-1].Done = true
			fmt.Println("-------------------------")
			log.Printf("[INFO] task completed: id=%d, text=%s", choice, List[choice-1].Text)
			fmt.Println("задача выполнена")
			fmt.Println("-------------------------")
		} else {
			fmt.Println("-------------------------")
			log.Println("[WARN] the task has already been marked")
			fmt.Println("задача уже отмечена")
			fmt.Println("-------------------------")
		}
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран не верный номер")
		fmt.Println("-------------------------")
	}
}

func UnmarkDoneTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	var choice int
	ShowTasks()
	fmt.Println("-------------------------")
	fmt.Println("выберите отметку какой задачи снять")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(List) {
		if List[choice-1].Done != false {
			List[choice-1].Done = false
			fmt.Println("-------------------------")
			log.Printf("[INFO] task unmarked: id=%d, text=%s", choice, List[choice-1].Text)
			fmt.Println("отметка убрана")
			fmt.Println("-------------------------")
		} else {
			fmt.Println("-------------------------")
			log.Printf("[WARN] there is no mark on this task")
			fmt.Println("отметки нет")
			fmt.Println("-------------------------")
		}
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран не верный номер")
		fmt.Println("-------------------------")
	}
}

func ShowTasks() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	for i, task := range List {
		if task.Done == true {
			fmt.Println(i+1, "[x]", task.Text)
		} else {
			fmt.Println(i+1, "[ ]", task.Text)
		}
	}
}
