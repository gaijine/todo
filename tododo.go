package main

import (

"fmt"

)

type Task struct{
	Text string
	Done bool
}

var list []Task

func menu(){
	var choice int
	for{	
	fmt.Println("-------------------------")
	fmt.Println("1. создать задачу")
	fmt.Println("2. вывести список задач")
	fmt.Println("3. пометить выполненной")
	fmt.Println("4. снять отметку")
	fmt.Println("0. выйти")
	fmt.Println("-------------------------")
	fmt.Scanln(&choice)
		
	var text string

		switch choice{		
	
		case 1:
			fmt.Println("введите текст задачи:")
			fmt.Scanln(&text)	
			task := Task{Text: text, Done: false}
			list = append(list, task)
		case 2:
			showTasks()
		case 3:
			markDoneTask()
		case 4:
			unmarkDoneTask()		
		case 0:
			return
		default:
			fmt.Println("ввели неверный номер")
		}
	
	}
}

func markDoneTask(){
	if len(list) == 0{
		fmt.Println("список задач пуст")
		return
	}

	var choice int
	showTasks()
	fmt.Println("выберите какую задачу пометить выполненной")
	fmt.Scanln(&choice)
	
	if choice >= 1 && choice <= len(list){
		if list[choice-1].Done != true{
		list[choice-1].Done = true
		fmt.Println("задача выполнена")
		}else{
		fmt.Println("задача уже отмечена")
		}
	}else{
	fmt.Println("выбран не верный номер")
	}
}

func unmarkDoneTask(){
	if len(list) == 0{
		fmt.Println("список задач пуст")
		return
	}

	var choice int
	showTasks()
	fmt.Println("выберите отметку какой задачи снять")
	fmt.Scanln(&choice)
	
	if choice >= 1 && choice <= len(list){
		if list[choice-1].Done != false{
		list[choice-1].Done = false
		fmt.Println("отметка убрана")
		}else{
		fmt.Println("отметки нет")
		}
	}else{
	fmt.Println("выбран не верный номер")
	}
}

func showTasks(){
	if len(list) == 0{
		fmt.Println("список задач пуст")
		return
	}

	for i, task := range list{
	if task.Done == true{
		fmt.Println(i+1,"[x]", task.Text)	
	}else{
		fmt.Println(i+1,"[ ]", task.Text)
	}
	}
}

func main(){

	menu()

}