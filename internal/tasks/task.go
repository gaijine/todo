package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AddTask() {
	reader := bufio.NewReader(os.Stdin) // Создаём новый объект "reader" типа bufio.Reader, который позволяет удобно считывать текст
	// из стандартного ввода (то есть с клавиатуры).
	// bufio.Reader помогает читать данные построчно или по частям, а не по одному символу.

	fmt.Println("-------------------------")
	fmt.Println("введите текст задачи:")
	fmt.Println("-------------------------")

	text, _ := reader.ReadString('\n') // С помощью reader читаем строку до символа '\n' (перенос строки), то есть до момента,
	// когда пользователь нажмёт Enter. ReadString возвращает две вещи: считанную строку и ошибку, но здесь ошибку игнорируем (_),
	// так как не обрабатываем её специально.
	text = strings.TrimSpace(text) // Функция strings.TrimSpace убирает все пробелы и символы переноса строки с начала и конца строки.
	// Это нужно, чтобы убрать лишние пробелы и особенно символ '\n', который остался в конце после нажатия Enter,
	// и получить "чистый" текст без пробелов вокруг.

	task := Task{Text: text, Done: false}
	List = append(List, task)
	fmt.Println("-------------------------")
	log.Printf("[INFO] task created: <%s>", task.Text)
	fmt.Printf("задача: <%s> успешно добавлена\n", task.Text)
}

func ShowTasksMenu() {
	fmt.Println("-------------------------")
	fmt.Println("выберите какие задачи показать:")
	fmt.Println("[1] вывести все задачи")
	fmt.Println("[2] только выполненные")
	fmt.Println("[3] только невыполненные")
	fmt.Println("-------------------------")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		ShowTasks()
	case 2:
		FilterTasks(true)
	case 3:
		FilterTasks(false)
	default:
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid menu number")
		fmt.Println("выбран неверный номер")
		fmt.Println("-------------------------")
	}
}

func ShowTasks() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		return
	}

	for i, task := range List {
		mark := "[ ]"
		if task.Done == true {
			mark = "[x]"
		}
		fmt.Println(i+1, mark, task.Text)
	}
}

func FilterTasks(showDone bool) {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		return
	}

	found := false

	if showDone == true {
		fmt.Println("-------------------------")
		fmt.Println("выполненные задачи:")
	} else {
		fmt.Println("-------------------------")
		fmt.Println("невыполненные задачи:")
	}

	for i, task := range List {
		if task.Done == showDone {
			mark := "[ ]"
			if task.Done {
				mark = "[x]"
			}
			fmt.Println(i+1, mark, task.Text)
			found = true
		}
	}

	if !found {
		fmt.Println("-------------------------")
		fmt.Println("задач с таким статусом нет")
	}
}

func MarkDoneTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
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
			log.Printf("[INFO] task completed: id=%d, text=<%s>", choice, List[choice-1].Text)
			fmt.Printf("задача: <%s> выполнена\n", List[choice-1].Text)
		} else {
			fmt.Println("-------------------------")
			log.Println("[WARN] the task has already been marked")
			fmt.Println("задача уже отмечена")
		}
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран не верный номер")
	}
}

func UnmarkDoneTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
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
			log.Printf("[INFO] task unmarked: id=%d, text=<%s>", choice, List[choice-1].Text)
			fmt.Printf("отметка выполнения задачи: <%s> снята\n", List[choice-1].Text)
		} else {
			fmt.Println("-------------------------")
			log.Printf("[WARN] there is no mark on this task")
			fmt.Println("отметки нет")
		}
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран не верный номер")
	}
}

func UpdateTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
		return
	}

	var choice int
	ShowTasks()
	fmt.Println("-------------------------")
	fmt.Println("выберите какую задачу изменить:")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(List) {

		reader := bufio.NewReader(os.Stdin) // создаём считыватель для ввода текста
		oldText := List[choice-1].Text

		fmt.Println("-------------------------")
		fmt.Println("введите новый текст задачи:")

		for {

			newText, _ := reader.ReadString('\n') // считываем строку с пробелами
			newText = strings.TrimSpace(newText)  // <-- убираем все лишние пробелы и перенос строки

			if newText == "" {
				fmt.Println("-------------------------")
				fmt.Printf("поле не заполнено, \nвведите текст повторно:\n")
			} else {
				List[choice-1].Text = newText
				fmt.Println("-------------------------")
				log.Printf("[INFO] task <%s> changed to <%s> successfully", oldText, newText)
				fmt.Printf("задача <%s> успешно изменена на <%s>\n", oldText, newText)
				return
			}
		}
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран неверный номер")
	}
}

func DeleteTask() {
	if len(List) == 0 {
		fmt.Println("-------------------------")
		log.Println("[WARN] task list is empty")
		fmt.Println("список задач пуст")
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
		log.Printf("[INFO] task deleted: id=%d, text=<%s>", choice, deletedTask.Text)
		fmt.Printf("задача: <%s> успешно удалена\n", deletedTask.Text)
	} else {
		fmt.Println("-------------------------")
		log.Println("[ERROR] invalid number task")
		fmt.Println("выбран неверный номер")
	}
}
