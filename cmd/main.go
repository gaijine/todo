package main

import (
	"log"
	"todo/internal/launcher"
	"todo/internal/tasks"
)

func main() {
	log.Println("[INFO] application started")
	tasks.LoadTasksFromFile()
	launcher.Menu()
}
