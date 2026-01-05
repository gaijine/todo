package main

import (
	"log"
	"todo/internal/launcher"
	"todo/internal/storage"
)

func main() {
	log.Println("[INFO] application started")
	storage.LoadTasksFromFile()
	launcher.Menu()
}
