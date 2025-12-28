package main

import (
	"fmt"
	"log"
)

type Task struct {
	Text string
	Done bool
}

var list []Task

func menu() {
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
			task := Task{Text: text, Done: false}
			list = append(list, task)
			log.Printf("[INFO] task created: %v", task.Text)
		case 2:
			showTasks()
		case 3:
			markDoneTask()
		case 4:
			unmarkDoneTask()
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

func markDoneTask() {
	if len(list) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	var choice int
	showTasks()
	fmt.Println("-------------------------")
	fmt.Println("выберите какую задачу пометить выполненной")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(list) {
		if list[choice-1].Done != true {
			list[choice-1].Done = true
			fmt.Println("-------------------------")
			log.Printf("[INFO] task completed: id=%v, text=%v", choice, list[choice-1].Text)
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

func unmarkDoneTask() {
	if len(list) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	var choice int
	showTasks()
	fmt.Println("-------------------------")
	fmt.Println("выберите отметку какой задачи снять")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(list) {
		if list[choice-1].Done != false {
			list[choice-1].Done = false
			fmt.Println("-------------------------")
			log.Printf("[INFO] unmarking completed successfully: id=%v, text=%v", choice, list[choice-1].Text)
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

func showTasks() {
	if len(list) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		fmt.Println("-------------------------")
		return
	}

	for i, task := range list {
		if task.Done == true {
			fmt.Println(i+1, "[x]", task.Text)
		} else {
			fmt.Println(i+1, "[ ]", task.Text)
		}
	}
}

func main() {
	log.Println("[INFO] application started")
	menu()

}
