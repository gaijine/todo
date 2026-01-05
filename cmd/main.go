package main

import (
	"log"
	"todo/internal/launcher"
	"todo/internal/logger"
	"todo/internal/storage"
)

func main() {
	logger.InitLogger()
	log.Println("[INFO] application started")
	storage.LoadTasksFromFile()
	launcher.Menu()
}
