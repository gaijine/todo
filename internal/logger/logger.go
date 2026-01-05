package logger

import (
	"log"
	"os"
)

func InitLogger() {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return
	}

	file, err := os.OpenFile("logs/app.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	log.SetOutput(file)
}
